[{
  name: .[].name,
  specs: [
    .[].specs[] |
      . + {scored: (.name | endswith("[Scored]"))} |
      .name |= rtrimstr(" [Scored]") |
      .name |= rtrimstr(" [Not Scored]")
  ] |
  group_by(.name) |
  map(reduce .[] as $item ({}; . * $item))
}] |
  unique |
  { 
    name: .[].name,
    specs: [
    .[].specs[] |
      . + {
        result: ([.results[].result] as $results |
          $results |
          map(
            if   . == "failed"    then 1
            elif . == "passed"    then 2
            elif . == "skipped"   then 3
            elif . == "pending"   then 4
            elif . == "timeout"   then 5
            elif . == "panicked"  then 6
            else                       7
            end
          ) | $results[index(min)]
        )
      }
    ] |
    sort_by(.name|scan("\\[\\d+\\.\\d+\\.\\d+\\]")|scan("\\d+\\.\\d+\\.\\d+")|split(".")|map(tonumber))
  } |
  . + {
    total:    [.specs[]] | length,
    passed:   [.specs[]] | map(select(.result == "passed"))   | length,
    failures: [.specs[]] | map(select(.result == "failed"))   | length,
    skipped:  [.specs[]] | map(select(.result == "skipped"))  | length,
    pending:  [.specs[]] | map(select(.result == "pending"))  | length,
    invalid:  [.specs[]] | map(select(.result == "invalid"))  | length,
    timeout:  [.specs[]] | map(select(.result == "timeout"))  | length,
    panicked: [.specs[]] | map(select(.result == "panicked")) | length,
    score:    {
      total_scored: [.specs[]] | map(select(.scored == true)) | length,
      scored_passed: [.specs[]] | map(select(.scored == true) | select(.result == "passed")) | length
    } | (.scored_passed * 100 / .total_scored) | round
  }
