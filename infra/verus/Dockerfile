FROM debian:buster-slim

# Install verusd.
RUN apt-get update                                                                                    && \
    apt-get install -y --no-install-recommends libgomp1 ca-certificates curl && \
    apt-get -y clean; \
    rm -rf /var/lib/apt/lists/* /tmp/* /var/tmp/*

RUN VERSION=$(curl -s https://api.github.com/repos/veruscoin/veruscoin/releases/latest | grep "tag_name" | awk '{print substr($2, 2, length($2)-3)}') && \
    PACKAGE=Verus-CLI-Linux-$VERSION-x86_64.tgz && \
    INNER_PACKAGE=Verus-CLI-Linux-$VERSION-x86_64.tar.gz && \
    curl -L -O  "https://github.com/VerusCoin/VerusCoin/releases/download/$VERSION/$PACKAGE" && \
    tar -xzvf "$PACKAGE" && \
    tar -xzvf "$INNER_PACKAGE" --directory /usr/bin/ --strip=1 && \
    fetch-params

COPY vrsctest.conf /root/.komodo/vrsctest/vrsctest.conf
COPY run.sh /root/run.sh
RUN chmod +x /root/run.sh

EXPOSE 45255

ENTRYPOINT ["./root/run.sh"]
