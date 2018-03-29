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
  [
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
