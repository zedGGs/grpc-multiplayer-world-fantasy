FROM envoyproxy/envoy:v1.25-latest
COPY ./service-envoy.yaml /etc/envoy/envoy.yaml
COPY ./localhost.pem /etc/localhost.pem
CMD /usr/local/bin/envoy -c /etc/envoy/envoy.yaml -l trace --log-path /tmp/envoy_info.log

<<<<<<< HEAD
EXPOSE 9901
=======
EXPOSE 9901
>>>>>>> a148ed9 (adding proto and build)
