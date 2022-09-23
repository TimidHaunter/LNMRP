## 查看信息

```shell
composer config -g -l
[repositories.packagist.org.type] composer
[repositories.packagist.org.url] https://repo.packagist.org
[process-timeout] 300
[use-include-path] false
[use-parent-dir] prompt
[preferred-install] dist
[notify-on-install] true
[github-protocols] [https, ssh]
[gitlab-protocol]
[vendor-dir] vendor (/www/yii2/vendor)
[bin-dir] {$vendor-dir}/bin (/www/yii2/vendor/bin)
[cache-dir] /tmp/composer/cache
[data-dir] /tmp/composer
[cache-files-dir] {$cache-dir}/files (/tmp/composer/cache/files)
[cache-repo-dir] {$cache-dir}/repo (/tmp/composer/cache/repo)
[cache-vcs-dir] {$cache-dir}/vcs (/tmp/composer/cache/vcs)
[cache-ttl] 15552000
[cache-files-ttl] 15552000
[cache-files-maxsize] 300MiB (314572800)
[cache-read-only] false
[bin-compat] auto
[discard-changes] false
[autoloader-suffix]
[sort-packages] false
[optimize-autoloader] false
[classmap-authoritative] false
[apcu-autoloader] false
[prepend-autoloader] true
[github-domains] [github.com]
[bitbucket-expose-hostname] true
[disable-tls] false
[secure-http] true
[cafile]
[capath]
[github-expose-hostname] true
[gitlab-domains] [gitlab.com]
[store-auths] prompt
[archive-format] tar
[archive-dir] .
[htaccess-protect] true
[use-github-api] true
[lock] true
[platform-check] php-only
[home] /tmp/composer
```



## 修改源

```shell
##阿里云
composer config -g repo.packagist composer https://mirrors.aliyun.com/composer/
##中国全量镜像
composer config -g repo.packagist composer https://packagist.phpcomposer.com
##腾讯云
composer config -g repos.packagist composer https://mirrors.cloud.tencent.com/composer/
```



## 解除镜像（还原）

```shell
composer config -g --unset repos.packagist
```

