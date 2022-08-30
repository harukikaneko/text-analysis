use itertools::Itertools;
use lindera::tokenizer::Token;

use super::{Noun, Nouns};

#[derive(Clone)]
pub struct Tokens<'a>(pub Vec<Token<'a>>);

impl Tokens<'_> {
    pub fn exclude_non_nouns(self) -> Nouns {
        Nouns(
            self.0
                .into_iter()
                .filter(|x| Detail(x.detail.clone()).is_nouns())
                .map(|nouns| Noun(nouns.text.into()))
                .collect_vec(),
        )
    }
}

#[derive(Debug, Clone, PartialEq)]
pub struct Detail(Vec<String>);

impl Detail {
    pub fn is_nouns(self) -> bool {
        self.0.contains(&"名詞".to_string()) || self.0.contains(&"カスタム名詞".to_string())
    }
}
