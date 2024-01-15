# BookCrudGolang

git merge --allow-unrelated-histories branch_to_merge

# step1 touch docker-compose.yml

        go mod init spy
# step2
        go get -u gorm.io/gorm
        go get -u gorm.io/driver/postgres
# step3
        go get -u github.com/gofiber/fiber/v2
# step4
        go get golang.org/x/crypto/bcrypt
        go get github.com/golang-jwt/jwt/v4