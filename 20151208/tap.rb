alphabet_nums = {}.tap do |h|
  ("a".. "z").each.with_index 1 do |a, i|
    h[a] = i
  end
end
puts alphabet_nums