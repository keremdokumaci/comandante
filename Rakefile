task :command_exists, [:command] do |_, args|
    abort "#{args.command} doesn't exists" if `command -v #{args.command} > /dev/null 2>&1 && echo $?`.chomp.empty?
  end
  
  task :has_hadolint do
    Rake::Task['command_exists'].invoke('hadolint')
  end
  
  task :has_golangcilint do
    Rake::Task['command_exists'].invoke('golangci-lint')
  end
  
  desc "run golangci-lint"
  task :lint => [:has_golangcilint] do
    system "LOG_LEVEL=error golangci-lint run"
  end