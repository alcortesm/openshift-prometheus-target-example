FROM scratch
COPY http-echo http-echo
CMD ["./http-echo"]
EXPOSE 8080
