#include <crypt.h>
#include <pwd.h>
#include <shadow.h>
#include <stdio.h>
#include <stdlib.h>
#include <string.h>
#include <sys/wait.h>
#include <unistd.h>

int checkSystemPassword(char *user, char *pass) {
  /*
    printf("Hello, %s!\n", user);
    printf("Hello, %s!\n", pass);
  */

  struct passwd *pw;

  pw = getpwnam(user);

  if (0 != strcmp(pw->pw_passwd, "x")) {

    return 2;

  } else {

    struct spwd *shadowEntry = getspnam(pw->pw_name);

    /*
        printf("%ld\n", shadowEntry->sp_expire);
        printf("%ld\n", shadowEntry->sp_flag);
        printf("%ld\n", shadowEntry->sp_inact);
        printf("%ld\n", shadowEntry->sp_lstchg);
        printf("%ld\n", shadowEntry->sp_max);
        printf("%ld\n", shadowEntry->sp_min);
        printf("%s\n", shadowEntry->sp_namp);
        printf("%ld\n", shadowEntry->sp_warn);
        printf("%s\n", shadowEntry->sp_pwdp);
     */

    char *p1 = crypt(pass, shadowEntry->sp_pwdp);

    if (strcmp(p1, shadowEntry->sp_pwdp) == 0) {
      return 1;
    } else {
      return 0;
    }
  }
}
