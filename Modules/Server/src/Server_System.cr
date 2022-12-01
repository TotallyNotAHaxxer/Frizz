
module Files 
    class FileWritersAndReaders
        def fwrite(filename, data)
            # Write to the file
            abort "Got error when writing to the file -> Missing path", 1 if !File.file? filename 
            filenametowrite = File.new filename, "w" # write mode
            filenametowrite.puts "#{data}"
            filenametowrite.close
        end 
        def freadtowritenew(inf, ouf)
            abort "Got error when writing to the file -> Missing path", 1 if !File.file? filename 
            wf = File.new ouf, "w"
            File.each_line inf do |line|
                wf.puts line
            end
            wf.close
        end
    end
end