package piscine

import "github.com/01-edu/z01"

func PrintCombN(n int) {
	if n >= 1 || n <= 9 {
		if n == 1 {
			for i := 48; i <= 57; i++ {
				z01.PrintRune(rune(i))
				if i == 57 {
					z01.PrintRune('\n')
				} else {
					z01.PrintRune(44)
					z01.PrintRune(32)
				}
			}
		} else if n == 2 {
			for i := 48; i <= 57; i++ {
				for j := 49; j <= 57; j++ {
					if i < j {
						z01.PrintRune(rune(i))
						z01.PrintRune(rune(j))
						if i == 56 && j == 57 {
							z01.PrintRune('\n')
						} else {
							z01.PrintRune(44)
							z01.PrintRune(32)
						}
					}
				}
			}
		} else if n == 3 {
			for i := 48; i <= 57; i++ {
				for j := 49; j <= 57; j++ {
					for k := 50; k <= 57; k++ {
						if i < j && j < k {
							z01.PrintRune(rune(i))
							z01.PrintRune(rune(j))
							z01.PrintRune(rune(k))
							if i == 55 && j == 56 && k == 57 {
								z01.PrintRune('\n')
							} else {
								z01.PrintRune(44)
								z01.PrintRune(32)
							}
						}
					}
				}
			}
		} else if n == 4 {
			for i := 48; i <= 57; i++ {
				for j := 49; j <= 57; j++ {
					for k := 50; k <= 57; k++ {
						for l := 51; l <= 57; l++ {
							if i < j && j < k && k < l {
								z01.PrintRune(rune(i))
								z01.PrintRune(rune(j))
								z01.PrintRune(rune(k))
								z01.PrintRune(rune(l))
								if i == 54 && j == 55 && k == 56 && l == 57 {
									z01.PrintRune('\n')
								} else {
									z01.PrintRune(44)
									z01.PrintRune(32)
								}
							}
						}
					}
				}
			}
		} else if n == 5 {
			for i := 48; i <= 57; i++ {
				for j := 49; j <= 57; j++ {
					for k := 50; k <= 57; k++ {
						for l := 51; l <= 57; l++ {
							for m := 52; m <= 57; m++ {
								if i < j && j < k && k < l && l < m {
									z01.PrintRune(rune(i))
									z01.PrintRune(rune(j))
									z01.PrintRune(rune(k))
									z01.PrintRune(rune(l))
									z01.PrintRune(rune(m))
									if i == 53 && j == 54 && k == 55 && l == 56 && m == 57 {
										z01.PrintRune('\n')
									} else {
										z01.PrintRune(44)
										z01.PrintRune(32)
									}
								}
							}
						}
					}
				}
			}
		} else if n == 6 {
			for i := 48; i <= 57; i++ {
				for j := 49; j <= 57; j++ {
					for k := 50; k <= 57; k++ {
						for l := 51; l <= 57; l++ {
							for m := 52; m <= 57; m++ {
								for o := 52; o <= 57; o++ {
									if i < j && j < k && k < l && l < m && m < o {
										z01.PrintRune(rune(i))
										z01.PrintRune(rune(j))
										z01.PrintRune(rune(k))
										z01.PrintRune(rune(l))
										z01.PrintRune(rune(m))
										z01.PrintRune(rune(o))
										if i == 52 && j == 53 && k == 54 && l == 55 && m == 56 && o == 57 {
											z01.PrintRune('\n')
										} else {
											z01.PrintRune(44)
											z01.PrintRune(32)
										}
									}
								}
							}
						}
					}
				}
			}
		} else if n == 7 {
			for i := 48; i <= 57; i++ {
				for j := 49; j <= 57; j++ {
					for k := 50; k <= 57; k++ {
						for l := 51; l <= 57; l++ {
							for m := 52; m <= 57; m++ {
								for o := 52; o <= 57; o++ {
									for p := 52; p <= 57; p++ {
										if i < j && j < k && k < l && l < m && m < o && o < p {
											z01.PrintRune(rune(i))
											z01.PrintRune(rune(j))
											z01.PrintRune(rune(k))
											z01.PrintRune(rune(l))
											z01.PrintRune(rune(m))
											z01.PrintRune(rune(o))
											z01.PrintRune(rune(p))
											if i == 51 && j == 52 && k == 53 && l == 54 && m == 55 && o == 56 && p == 57 {
												z01.PrintRune('\n')
											} else {
												z01.PrintRune(44)
												z01.PrintRune(32)
											}
										}
									}
								}
							}
						}
					}
				}
			}
		} else if n == 8 {
			for i := 48; i <= 57; i++ {
				for j := 49; j <= 57; j++ {
					for k := 50; k <= 57; k++ {
						for l := 51; l <= 57; l++ {
							for m := 52; m <= 57; m++ {
								for o := 52; o <= 57; o++ {
									for p := 52; p <= 57; p++ {
										for q := 52; q <= 57; q++ {
											if i < j && j < k && k < l && l < m && m < o && o < p && p < q {
												z01.PrintRune(rune(i))
												z01.PrintRune(rune(j))
												z01.PrintRune(rune(k))
												z01.PrintRune(rune(l))
												z01.PrintRune(rune(m))
												z01.PrintRune(rune(o))
												z01.PrintRune(rune(p))
												z01.PrintRune(rune(q))
												if i == 50 && j == 51 && k == 52 && l == 53 && m == 54 && o == 55 && p == 56 && q == 57 {
													z01.PrintRune('\n')
												} else {
													z01.PrintRune(44)
													z01.PrintRune(32)
												}
											}
										}
									}
								}
							}
						}
					}
				}
			}
		} else if n == 9 {
			for i := 48; i <= 57; i++ {
				for j := 49; j <= 57; j++ {
					for k := 50; k <= 57; k++ {
						for l := 51; l <= 57; l++ {
							for m := 52; m <= 57; m++ {
								for o := 52; o <= 57; o++ {
									for p := 52; p <= 57; p++ {
										for q := 52; q <= 57; q++ {
											for r := 52; r <= 57; r++ {
												if i < j && j < k && k < l && l < m && m < o && o < p && p < q && q < r {
													z01.PrintRune(rune(i))
													z01.PrintRune(rune(j))
													z01.PrintRune(rune(k))
													z01.PrintRune(rune(l))
													z01.PrintRune(rune(m))
													z01.PrintRune(rune(o))
													z01.PrintRune(rune(p))
													z01.PrintRune(rune(q))
													z01.PrintRune(rune(r))
													if i == 49 && j == 50 && k == 51 && l == 52 && m == 53 && o == 54 && p == 55 && q == 56 && r == 57 {
														z01.PrintRune('\n')
													} else {
														z01.PrintRune(44)
														z01.PrintRune(32)
													}
												}
											}
										}
									}
								}
							}
						}
					}
				}
			}
		}
	}
}
