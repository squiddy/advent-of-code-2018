import math
import parseutils
import sets

proc readFrequencyDrifts(): seq[int] =
    result = @[]
    var file: File
    var line: TaintedString
    var value: int
    discard system.open(file, "../input.txt")
    while system.readLine(file, line):
        if parseutils.parseInt(line, value) > 0:
            result.add(value)
    
    return result

proc getFinalFrequency(drifts: seq[int]): int = math.sum(drifts)

proc getFirstDuplicateFrequency(drifts: seq[int]): int =
    var frequency = 0
    var seen = initSet[int]();

    while true:
        for drift in drifts:
            frequency += drift
            if frequency in seen:
                return frequency
            else:
                seen.incl(frequency)

var drifts = readFrequencyDrifts()
echo getFinalFrequency(drifts)
echo getFirstDuplicateFrequency(drifts)