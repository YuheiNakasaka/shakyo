class Ruby
  def method_missing(method, *args)
    puts "#{method}が呼ばれました"
  end
end

r = Ruby.new
r.hello