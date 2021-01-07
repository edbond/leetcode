# @param {String} s
# @return {Integer}
require 'set'

def length_of_longest_substring(s)
    return 0 if s.to_s == ""
    
    i, j = 0, 0
    longest = ""
    chars = Set.new()
    
    loop do
        break if i > s.size || j > s.size
        
        # puts "i #{i}, j #{j}, s[j] #{s[j]}, chars #{chars}, longest #{longest}"
        
        if chars.include?(s[j])
            chars.delete(s[i])
            i += 1            
            next
        end
        
        chars << s[j]
        if j + 1 - i > longest.size
            longest = s[i..j]
            # puts "new longest #{longest}"
        end
        j += 1
    end
    
    longest.size
end