# @param {String} s
# @return {Boolean}
CLOSE = {
    ")" => "(",
    "}" => "{",
    "]" => "[",
}.freeze
OPEN = [ "(", "[", "{" ].freeze

def is_valid(s)
    r = s.chars.each_with_object([]) do |i,acc|
        if OPEN.include?(i)
            acc << i
            next
        else
            if acc[-1] == CLOSE.fetch(i)
                acc.pop
                next
            else
                return false
            end
        end
    end
    r.empty?
end