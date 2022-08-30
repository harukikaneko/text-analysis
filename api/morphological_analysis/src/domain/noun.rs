use itertools::Itertools;
use serde::Serialize;

#[derive(Debug, Clone, PartialEq)]
pub struct Nouns(pub Vec<Noun>);

impl Nouns {
    pub fn aggregate_group_by_word(self) -> Vec<CountByNoun> {
        self.0
            .into_iter()
            .into_group_map_by(|word| word.to_owned())
            .into_iter()
            .map(|(key, values)| -> CountByNoun {
                CountByNoun {
                    noun: key,
                    counts: values.len(),
                }
            })
            .collect_vec()
    }
}

#[derive(Debug, Serialize, Clone, PartialEq, Hash, Eq)]
pub struct Noun(pub String);

#[derive(Debug, Serialize, Clone, PartialEq)]
pub struct CountByNoun {
    pub noun: Noun,
    pub counts: usize,
}
