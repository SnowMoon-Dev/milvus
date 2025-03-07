# Licensed to the LF AI & Data foundation under one
# or more contributor license agreements. See the NOTICE file
# distributed with this work for additional information
# regarding copyright ownership. The ASF licenses this file
# to you under the Apache License, Version 2.0 (the
# "License"); you may not use this file except in compliance
# with the License. You may obtain a copy of the License at
#
#     http://www.apache.org/licenses/LICENSE-2.0
#
# Unless required by applicable law or agreed to in writing, software
# distributed under the License is distributed on an "AS IS" BASIS,
# WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
# See the License for the specific language governing permissions and
# limitations under the License.

if [ "$(expr substr $(uname -s) 1 5)" == "Linux" ]; then
	LIBJEMALLOC=$PWD/internal/core/output/lib/libjemalloc.so
	if test -f "$LIBJEMALLOC"; then
		#echo "Found $LIBJEMALLOC"
		export LD_PRELOAD="$LIBJEMALLOC"
		echo export LD_PRELOAD="$LIBJEMALLOC"
	else
		echo "WARN: Cannot find $LIBJEMALLOC"	
	fi
fi	

echo "Starting standalone..."
nohup ./bin/milvus run standalone > /tmp/standalone.log 2>&1 &
