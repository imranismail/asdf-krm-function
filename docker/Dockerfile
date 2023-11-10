FROM public.ecr.aws/ubuntu/ubuntu:23.04_stable

SHELL ["/usr/bin/env", "-S", "bash", "-cl"]

RUN apt-get update -q && apt-get install -y git curl python3 python3-pip && \
    apt-get clean && \
    rm -rf /var/lib/apt/lists/*

RUN adduser --uid 999 --shell /bin/bash --home /home/argocd --disabled-password argocd
ENV PATH="${PATH}:/asdf/.asdf/shims:/asdf/.asdf/bin"

USER argocd
WORKDIR /home/argocd

RUN git clone --depth 1 https://github.com/asdf-vm/asdf.git $HOME/.asdf && \
    echo '. $HOME/.asdf/asdf.sh' >> $HOME/.bashrc && \
    echo '. $HOME/.asdf/asdf.sh' >> $HOME/.profile

RUN asdf plugin add kustomize && \
    asdf plugin add helm && \
    asdf plugin add krm-function https://github.com/imranismail/asdf-krm-function.git && \
    asdf install kustomize latest && \
    asdf global kustomize latest && \
    asdf install helm latest && \
    asdf global helm latest

COPY plugin /usr/local/bin/
COPY plugin.yaml ./cmp-server/config/
