"""Functions for creating, transforming, and adding prefixes to strings."""


def add_prefix_un(word):
    """Take the given word and add the 'un' prefix.

    Parameters:
        word (str): The root word.

    Returns:
        str: Root word prepended with 'un'.
    """

    return "un" + word


def make_word_groups(vocab_words):
    """Transform a list containing a prefix and words.

    Parameters:
        vocab_words (list[str]): Vocabulary words with prefix at first index.

    Returns:
        str: Prefix followed by vocabulary words with prefix applied.

    This function takes a `vocab_words` list of strings and returns a string
    with the prefix and the words with prefix applied, separated by ' :: '.

    Examples:
        >>> list('en', 'close', 'joy', 'lighten')
        'en :: enclose :: enjoy :: enlighten'.

    """
    prefixed_words = []
    words_to_prefix = vocab_words[1:]
    for _, item in enumerate(words_to_prefix):
        prefixed_words.append(vocab_words[0] + item)
    final_word_list = []
    final_word_list.append(vocab_words[0])
    for _, item in enumerate(prefixed_words):
        final_word_list.append(item)
    return " :: ".join(final_word_list)


def remove_suffix_ness(word):
    """Remove the suffix from the word while keeping spelling in mind.

    Parameters:
        word (str): Word to remove suffix from.

    Returns:
        str: Word with suffix removed & spelling adjusted.

    Examples:
        >>> remove_suffix_ness('heaviness')
        'heavy'

        >>> remove_suffix_ness('sadness')
        'sad'

    """
    if word.endswith("iness"):
        replaced = str.replace(word, "iness", "y")
        return replaced
     
    elif word.endswith("ness"):
        replaced = str.replace(word, "ness", "")
        return replaced

    
def adjective_to_verb(sentence, index):
    """Change the adjective within the sentence to a verb.

    Parameters:
        sentence (str): The word used in a sentence as an adjective.
        index (int): Index of the adjective to remove and transform.

    Returns:
        str: The extracted adjective in verb form.

    Examples:
        >>> adjective_to_verb('It got dark as the sun set.', 2)
        'darken'

        >>> adjective_to_verb('The ink stains her fingers black.', -1)
        'blacken'

    """
    # remove punctuation
    sentence_no_punct = str.replace(sentence, ".", "")
    sentence_no_punct = str.replace(sentence_no_punct, ",", "")
    sentence_no_punct = str.replace(sentence_no_punct, "?", "")
    sentence_no_punct = str.replace(sentence_no_punct, "!", "")    
    sentence_array = str.split(sentence_no_punct, " ")
    verbed_adjective = sentence_array[index] + "en"
    return verbed_adjective
