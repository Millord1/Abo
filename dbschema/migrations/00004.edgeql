CREATE MIGRATION m16tdki4dtbfhxm53ppgts3pqnhhht2bsewejfatgr7l3ukk3yuusa
    ONTO m1mciecc3zry3v7fwaitqvadvemufxwgs4xjxa4wpk2ahkfh3x6uqa
{
  ALTER TYPE default::Asset {
      ALTER PROPERTY description {
          SET REQUIRED USING ('null');
      };
  };
  ALTER TYPE default::Asset {
      ALTER PROPERTY metadata {
          SET REQUIRED USING ('null');
      };
  };
};
