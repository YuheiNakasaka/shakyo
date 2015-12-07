class DSL
  attr_reader :options

  DEFAULT = {
    foo: 'default foo',
    boolean_value: false
  }.map { |k, v| [k, v.freeze] }.to_h.freeze

  def self.load(options, path = nil)
    new(options).tap do |obj|
      obj._load_from(path)
    end
  end

  def initialize(options)
    @options = DEFAULT.dup
    @options.merge!(options)
  end

  def _load_from(path)
    instance_eval(File.read(path), path) if path
  end

  def foo(str)
    @options[:foo] = str
  end

  def block_test(&block)
    block.call
  end

  def boolean_test
    @options[:boolean_value] = true
  end
end

p DSL.load({}, 'config.rb')