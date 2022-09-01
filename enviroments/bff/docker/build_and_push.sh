USAGE="[-t <TAG>]"

while getopts t:h opt; do
  case $opt in
      t) tag=${OPTARG}
         ;;
      *) usage
         ;;
  esac
done
shift $((OPTIND - 1))

IMAGE_NAME_AND_TAG=niku100g/text-analysis-bff:${tag:-"latest"}

DOCKER_BUILDKIT=1 docker build -t $IMAGE_NAME_AND_TAG -f ./Dockerfile ../../../bff
docker push $IMAGE_NAME_AND_TAG
