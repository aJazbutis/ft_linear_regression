FROM ubuntu
RUN apt update && apt upgrade
RUN apt install -y golang-go
RUN apt install -y gnuplot-data libc6 libcairo2 libedit2 libgcc-s1 libgcc-s1 libgd3 libglib2.0-0 liblua5.3-0 libpango-1.0-0 libpangocairo-1.0-0 libstdc++6 libwxgtk3.0-gtk3-0v5 libx11-6
RUN apt install -y gnuplot-x11
WORKDIR /goproject
CMD ["sleep", "infinity"]
