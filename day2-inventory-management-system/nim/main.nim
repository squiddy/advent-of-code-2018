import sequtils
import tables

var boxIds = toSeq("../input.txt".lines)

# part 1

proc countLetters(value: string): CountTable[char] =
    result = initCountTable[char]()
    for character in value:
        result.inc(character)

var boxesWithDoubleLetters = 0
var boxesWithTripleLetters = 0
for id in boxIds:
    var counts = countLetters(id)
    if any(toSeq(counts.values), proc (x: int): bool = x == 2):
        boxesWithDoubleLetters += 1
    if any(toSeq(counts.values), proc (x: int): bool = x == 3):
        boxesWithTripleLetters += 1

echo boxesWithDoubleLetters * boxesWithTripleLetters

# part 2

for idx, idA in boxIds:
    for idB in boxIds[idx..boxIds.len - 1]:
        var differences = 0
        for pair in zip(idA, idB):
            if pair.a != pair.b:
                differences += 1

        if differences == 1:
            var common: seq[char] = @[]
            for pair in zip(idA, idB):
                if pair.a == pair.b:
                    common.add(pair.a)
            echo cast[string](common)