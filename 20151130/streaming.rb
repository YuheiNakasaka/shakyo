require 'twitter'

client = Twitter::Streaming::Client.new do |config|
  config.consumer_key = ""
  config.consumer_secret = ""
  config.access_token = ""
  config.access_token_secret = ""
end

# client.sample do |obj|
#   puts obj.text if obj.is_a?(Twitter::Tweet)
# end

topics = ["猫"]
client.filter(track: topics.join(',')) do |obj|
  puts "#{obj.user.screen_name} #{obj.text}" if obj.is_a?(Twitter::Tweet)
end