

Install rr
https://rr-project.org/

```
sudo dnf install rr -y
```

Kernel changes for RR
```
sudo sysctl kernel.perf_event_paranoid=-1
sudo sysctl kernel.kptr_restrict=0
```


Run RR with `openshift-install`
```
rr openshift-install create cluster --log-level debug
```

after install failed run

```
rr pack
```

**WARNING** this will be very large and should be encrypted
Tar the e.g. `/home/<user>/.local/share/rr` directory

```
tar -czvf openshift-debug.tar.gz ${HOME}/.local/share/rr
```

Grab the `debug.log` from the directory openshift-install was executed from.


#### engineering

```
dlv replay --backend=rr --headless --listen=:2345 --api-version=2 --accept-multiclient /home/jcallen/.local/share/rr/openshift-install-0
```
then use goland to debug

Watch out for CPU compatiablity like avx512. VMC has avx512 though.




