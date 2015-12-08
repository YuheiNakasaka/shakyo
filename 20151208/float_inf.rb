to ||= Float::INFINITY
[1,2,3].each do |x|
  if x < to
    puts x
  end
end