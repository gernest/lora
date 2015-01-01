
task :clean do
  sh %{ compass clean}
end

task :compass=>:clean do
  sh %{ bundle exec compass watch}
end

task :s do
  sh %{bee run}
end

task :selenium do
  sh %{export DISPLAY=":99" && java -jar /usr/bin/selenium/selenium-server.jar }
end

task :test do
  sh %{ ginkgo -r }
end

task :test_dev do
  sh %{ ginkgo watch -r  -depth=0}
end