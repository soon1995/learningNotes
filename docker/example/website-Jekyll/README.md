# On-demand website using the Jekyll framework

1. Create the Jekyll base image and the Apache image (once-off)

    - see JekyllBaseImage dir -> `docker build -t jamtur01/jekyll .`

    - see ApacheBaseImage dir -> `docker build -t jamtur01/apache .`

2. Create a container from our Jekyll image that holds our website source mounted via a volume.

    - ```bash
      docker run -v /home/topsal/Repos/learningNotes/docker/example/website-Jekyll/sourceCode/james_blog:/data/ \
      --name james_blog jamtur01/jekyll
      ```

    - `docker inspect -f "{{ range .Mounts }}{{.}}{{end}}" james_blog` 

3. Create a Docker container from our Apache image that uses the volume containing
the compiled site and serve that out.

    - share volume with james_blog

    - ```bash
      docker run -d -P --volumes-from james_blog jamtur01/apache
      ```

    - you can create backups of the volumes if you are a little worried about accidentally deleting your volume

      - ```bash
        docker run --rm --volumes-from james_blog \
        -v $(pwd):/backup ubuntu \
        tar cvf /backup/james_blog_backup.tar /var/www/html
        ```

4. Rinse and repeat as the site needs to be updated.
    
    - make changes in sourceCode

    - `docker start james_blog`

You could consider this a simple way to create multiple hosted website instances.

## Extending our Jekyll website example

> Run multiple Apache containers, all which use the same volume from the james_blog container. Put a load
balancer in front of it, and we have a web cluster

> Build a further image that cloned or copied a user-provided source (e.g. a git clone) into a volume. Mount
this volume into a container created from our jamtur01/jekyll image. This would make the solution portable
and generic and would not require any local source on a host.

> With the previous expansion, you could easily build a web front end for our service that built and deployed sites automatically
from a specified source. Then you would have your variant of GitHub Pages.
