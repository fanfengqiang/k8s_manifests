ldapsearch -x -H ldap:// -b dc=5ik8s,dc=com -D "cn=admin,dc=5ik8s,dc=com" -w admin
ldapadd -x -D "cn=admin,dc=5ik8s,dc=com" -w admin -f ./groups.ldif -H ldap:/// 
ldapadd -x -D "cn=admin,dc=5ik8s,dc=com" -w admin -f ./users.ldif -H ldap:/// 
