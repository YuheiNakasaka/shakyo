user_ids = Enumerator.new do |y|
  open("candidates.txt") do |f|
    f.each_line do |line|
      next if line.start_with?("#")
      user_id = line.chomp
      y << user_id
    end
  end
end

user_ids.with_index 1 do |user_id, index|
  puts "#{index} #{user_id}"
end