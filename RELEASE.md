# Release

```bash
# Build Helm Charts at the master branch
helm package helm-charts/
git checkout gh-pages
cp gpu-operator*.tgz ./charts/

# Update the index.yml
helm repo index . --url https://rocm.github.io/gpu-operator

# Release
git add ./charts/*.tgz
git add index.yaml
git commit -m 'Release version XXX'

# deploy the new GitHub page
git push 
```
