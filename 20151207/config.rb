# new(options).tap do |obj|
#   obj._load_from(path)
# end
# なので
# config.rbのコンテキストはDSLクラスのインスタンスオブジェクト

foo 'test'

block_test do
  puts 'blockも渡せる'
end

boolean_test