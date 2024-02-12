FROM envoyproxy/envoy:v1.25-latest
COPY ./local-envoy.yaml /etc/envoy/envoy.yaml
COPY ./localhost.pem /etc/localhost.pem
CMD /usr/local/bin/envoy -c /etc/envoy/envoy.yaml -l trace --log-path /tmp/envoy_info.log

<<<<<<< HEAD
EXPOSE 9911
=======
EXPOSE 9911
>>>>>>> a148ed9 (adding proto and build)
