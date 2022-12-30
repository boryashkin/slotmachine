# The Math

Reels:
first second third
| reel1 | reel2 | reel3 |
|---|---|---|
| 0 | 0 | 0 |
| 1 | 1 | 1 |
| 2 | 2 | 2 |
| . | . | . |
| 9 | 9 | 9 |

total combinations: 10 * 10 * 10 = 1000

Probabilities:

(k - certain value of a reel, x - any value)
| reels | probability | % |
|---|---|---|
| x x x | 1/1000 = 0.001 | 0.1% |
| k k k | 10/1000 = 0.01 | 1% |
| k k x | 100/1000 = 0.1 | 10% |

Inputs

1 per spin

Payouts:

(sum of payouts should match RTP: if total combination amount is 1000 and input = 1 and RTP is 95%, sum of payouts should be 950)

See SlotSetPayout struct
