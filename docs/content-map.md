# Cipher-Quest Content Map

Campaign index. Dialogue and payload variants live in `internal/content/scripts_ch*.go`. Wire a new cipher with:

```go
Cipher: YourType{...}, Content: pickIntercepts(YourType{...}, thatLevelIntercepts)
```

Only Caesar and Atbash are implemented. Every other level is a menu stub plus a ready `[]interceptStage` script.

Unlocks are sequential: finish a level to open the next, including across chapters.

## Chapter 1 — Substitution Ciphers: Caesar's March

**Setting:** Your office. Roman signals desk.
**Plot:** Caesar is moving on the city. Muster → new route → target → launch.

| ID | Level | Scripts | Expected cipher | Key hint |
| --- | --- | --- | --- | --- |
| `caesar` | Caesar Cipher | `caesarIntercepts` | `CaesarCipher{Shift: 3}` (implemented) | loves **3** |
| `atbash` | Atbash Cipher | `atbashIntercepts` | `AtbashCipher{}` (implemented) | **mirrors** |
| `keyword` | Keyword Cipher | `keywordIntercepts` | `KeywordCipher`, key `JULIUS` | mixed from **his own name** |
| `affine` | Affine Cipher | `affineIntercepts` | `AffineCipher`, keys `5, 8` | **two** lucky numbers |

## Chapter 2 — Transposition Ciphers: The Scrambled Galleon

**Setting:** Admiralty cabin. Butts is on loan and seasick.
**Plot:** Captain Widdershins and Patch scramble every letter. Find the galleon → stop the silver raid → beat them to the stash.

| ID | Level | Scripts | Expected cipher | Key hint |
| --- | --- | --- | --- | --- |
| `scytale` | Scytale Cipher | `scytaleIntercepts` | `ScytaleCipher`, staff width | wrapped around a **belaying pin** |
| `columnar` | Columnar Transposition | `columnarIntercepts` | `ColumnarCipher`, keyword / column count | **ledger columns** |
| `rail_fence` | Rail Fence Cipher | `railFenceIntercepts` | `RailFenceCipher`, rail count | zigzag **ratlines** |
| `route` | Route Cipher | `routeIntercepts` | `RouteCipher`, spiral / deck plan | around the **deck from the helm** |

## Chapter 3 — Polyalphabetic Ciphers: Amore Infinito

**Setting:** Carnival in Venice. A box at a flop opera.
**Plot:** Maestro Recitativo sells secrets in *Amore Infinito*. Buyer → drop → masquerade → double-cross.

| ID | Level | Scripts | Expected cipher | Key hint |
| --- | --- | --- | --- | --- |
| `vigenere` | Vigenère Cipher | `vigenereIntercepts` | `VigenereCipher`, keyword `AMORE` | love-duet word that **never ends** |
| `gronsfeld` | Gronsfeld Cipher | `gronsfeldIntercepts` | `GronsfeldCipher`, digits `31415` | **box-office receipts** |
| `autokey` | Autokey Cipher | `autokeyIntercepts` | `AutokeyCipher`, primer `AMORE` | he got **paranoid** |
| `beaufort` | Beaufort Cipher | `beaufortIntercepts` | `BeaufortCipher`, reciprocal | sang the aria **backwards** |

## Chapter 4 — Modern Ciphers: The Boinkers File

**Setting:** A 1917 wireless hut.
**Plot:** Lieutenant Boinkers invents four unbreakable systems a month. Battery → new square → fake pad → he mails the real key with the telegram.

| ID | Level | Scripts | Expected cipher | Key hint |
| --- | --- | --- | --- | --- |
| `playfair` | Playfair Cipher | `playfairIntercepts` | `PlayfairCipher`, 5×5, key `BOINK` | letters travel in **pairs** |
| `bifid` | Bifid Cipher | `bifidIntercepts` | `BifidCipher`, period 5 | he **fractionated** the alphabet |
| `boinkers` | Boinkers Cipher | `boinkersIntercepts` | repeating XOR / fake OTP, key `BOINK` | a pad of **five letters** |
| `vernam` | Vernam Cipher | `vernamIntercepts` | `VernamCipher`, true OTP | key in the **same envelope** |
