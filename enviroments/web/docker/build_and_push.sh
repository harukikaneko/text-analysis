USAGE="[-t <TAG>]"
USAGE="[-a <ARG>]"

while getopts a:t:h opt; do
  case $opt in
      a) arg=${OPTARG}
         echo "$OPTARG";;
      t) tag=${OPTARG}
         echo "$OPTARG";;
      *) usage
         ;;
  esac
done
shift $((OPTIND - 1))

IMAGE_NAME_AND_TAG=niku100g/text-analysis-web:${tag:-"latest"}
ARG=${arg:-".env.development"}

DOCKER_BUILDKIT=1 docker build -f ./Dockerfile ../../../web --build-arg ENV_FILE_NAME=$ARG -t $IMAGE_NAME_AND_TAG
docker push $IMAGE_NAME_AND_TAG
