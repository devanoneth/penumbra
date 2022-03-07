use std::cell::Cell;

use crate::{Elems, GetHash, Hash, Height, Three};

#[derive(Debug, Clone, Eq, PartialEq)]
pub(crate) struct Active<Focus: crate::Active> {
    focus: Focus,
    siblings: Three<Result<Focus::Complete, Hash>>,
    // TODO: replace this with space-saving `Cell<OptionHash>`?
    hash: Cell<Option<Hash>>,
}

impl<Focus: crate::Active> Active<Focus> {
    pub(crate) fn from_parts(siblings: Three<Result<Focus::Complete, Hash>>, focus: Focus) -> Self
    where
        Focus: crate::Active + GetHash,
    {
        Self {
            hash: Cell::new(None),
            siblings,
            focus,
        }
    }
}

fn hash_active<Focus: crate::Active + GetHash>(
    siblings: &Three<Result<Focus::Complete, Hash>>,
    focus: &Focus,
) -> Hash {
    // Get the correct padding hash for this height
    let padding = Hash::padding();

    // Get the hashes of all the `Result<T, Hash>` in the array, hashing `T` as necessary
    #[inline]
    fn hashes_of_all<T: GetHash, const N: usize>(full: [&Result<T, Hash>; N]) -> [Hash; N] {
        full.map(|result| {
            result
                .as_ref()
                .map(|complete| complete.hash())
                .unwrap_or_else(|hash| *hash)
        })
    }

    // Get the four elements of this segment, *in order*, and extract their hashes
    let (a, b, c, d) = match siblings.elems() {
        Elems::_0([]) => {
            let a = focus.hash();
            let [b, c, d] = [padding, padding, padding];
            (a, b, c, d)
        }
        Elems::_1(full) => {
            let [a] = hashes_of_all(full);
            let b = focus.hash();
            let [c, d] = [padding, padding];
            (a, b, c, d)
        }
        Elems::_2(full) => {
            let [a, b] = hashes_of_all(full);
            let c = focus.hash();
            let [d] = [padding];
            (a, b, c, d)
        }
        Elems::_3(full) => {
            let [a, b, c] = hashes_of_all(full);
            let d = focus.hash();
            (a, b, c, d)
        }
    };

    Hash::node(Focus::HEIGHT + 1, a, b, c, d)
}

impl<Focus: crate::Active> Height for Active<Focus>
where
    Focus: Height,
{
    const HEIGHT: usize = if Focus::HEIGHT == <Focus as crate::Active>::Complete::HEIGHT {
        Focus::HEIGHT + 1
    } else {
        panic!("`Sibling::HEIGHT` does not match `Focus::HEIGHT` in `Segment`: check for improper depth types")
    };
}

impl<Focus: crate::Active> GetHash for Active<Focus> {
    fn hash(&self) -> Hash {
        match self.hash.get() {
            Some(hash) => hash,
            None => {
                let hash = hash_active(&self.siblings, &self.focus);
                self.hash.set(Some(hash));
                hash
            }
        }
    }
}

impl<Focus> crate::Active for Active<Focus>
where
    Focus: crate::Active + GetHash,
{
    type Item = Focus::Item;
    type Complete = super::Complete<Focus::Complete>;

    #[inline]
    fn singleton(item: Self::Item) -> Self {
        let focus = Focus::singleton(item);
        let siblings = Three::new();
        Self::from_parts(siblings, focus)
    }

    #[inline]
    fn complete(self) -> Result<Self::Complete, Hash> {
        super::Complete::from_siblings_and_focus_or_else_hash(self.siblings, self.focus.complete())
    }

    #[inline]
    fn alter<T>(&mut self, f: impl FnOnce(&mut Self::Item) -> T) -> Option<T> {
        let result = self.focus.alter(f);
        self.hash.set(None); // the hash could have changed, so clear the cache
        result
    }

    #[inline]
    fn insert(self, item: Self::Item) -> Result<Self, (Self::Item, Result<Self::Complete, Hash>)> {
        match self.focus.insert(item) {
            // We successfully inserted at the focus, so siblings don't need to be changed
            Ok(focus) => Ok(Self::from_parts(self.siblings, focus)),

            // We couldn't insert at the focus because it was full, so we need to move our path
            // rightwards and insert into a newly created focus
            Err((item, sibling)) => match self.siblings.push(sibling) {
                // We had enough room to add another sibling, so we set our focus to a new focus
                // containing only the item we couldn't previously insert
                Ok(siblings) => Ok(Self::from_parts(siblings, Focus::singleton(item))),

                // We didn't have enough room to add another sibling, so we return a complete node
                // as a carry, to be propagated up above us and added to some ancestor segment's
                // siblings, along with the item we couldn't insert
                Err(complete) => {
                    Err((
                        item,
                        // Implicitly, we only hash the children together when we're pruning them
                        // (because otherwise we would lose that informtion); if at least one child
                        // and its sibling hashes/subtrees is preserved in a `Complete` node, then
                        // we defer calculating the node hash until looking up an authentication path
                        super::Complete::from_children_or_else_hash(complete).map(|node| {
                            if let Some(hash) = self.hash.get() {
                                // This is okay because `complete` is guaranteed to have the same elements in
                                // the same order as `siblings + [focus]`:
                                node.set_hash_unchecked(hash)
                            }
                            node
                        }),
                    ))
                }
            },
        }
    }
}
