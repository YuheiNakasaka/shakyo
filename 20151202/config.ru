class NullEndPoint
  def call(env)
  end
end

use Rack::ETag
use Rack::Deflater
use Rack::Static, urls: [''], root: 'public', index: 'index.html'
run NullEndPoint.new
