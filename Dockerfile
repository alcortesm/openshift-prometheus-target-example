FROM scratch
COPY http-echo http-echo
EXPOSE 8080
USER 10001
CMD ["./http-echo"]
