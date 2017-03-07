function anylize_dna_sequence(sequence) {
  return valid.match(/([CGAT]{3}){1,}/g)
}
anylize_dna_sequence("ATATTGGTGTTCATGTGCGCGGGGCCGACGAGCTACTGGCAGAACCACGAGGACAAGAGGTGA")
anylize_dna_sequence("FAIL")
anylize_dna_sequence("Alanine")
