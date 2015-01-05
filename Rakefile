# Copyright 2015 Geofrey Ernest a.k.a gernest, All Rights Reserved.
#
# Licensed under the Apache License, Version 2.0 (the "License"): you may
#  not use this file except in compliance with the License. You may obtain
#  a copy of the License at
#
#      http://www.apache.org/licenses/LICENSE-2.0
#
# Unless required by applicable law or agreed to in writing, software
# distributed under the License is distributed on an "AS IS" BASIS, WITHOUT
# WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied. See the
# License for the specific language governing permissions and limitations
# under the License.


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