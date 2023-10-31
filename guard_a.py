import psutil
import os
import signal
 
print("----------------------------- show all processes info --------------------------------")
# show processes info
pids = psutil.pids()
for pid in pids:
    p = psutil.Process(pid)
    # get process name according to pid
    process_name = p.name()
    
    print("Process name is: %s, pid is: %s" %(process_name, pid))
 
print("----------------------------- kill specific process --------------------------------")
pids = psutil.pids()
while True:
    is_exist = False
    for pid in pids:
        p = psutil.Process(pid)
        # get process name according to pid
        process_name = p.name()
        # kill process "sleep_test1"
        if process_name == 'tree_hole_s':
            is_exist = True
            break
    if is_exist == False:
        