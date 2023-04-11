CREATE MIGRATION m14rruqf2oruky3gs2ygcqljhipf5fdvrji6syior6upfe5pwmkjea
    ONTO initial
{
  CREATE FUTURE nonrecursive_access_policies;
  CREATE TYPE default::Asset {
      CREATE REQUIRED PROPERTY assetId -> std::str;
      CREATE PROPERTY description -> std::str;
      CREATE REQUIRED PROPERTY imgUrl -> std::str;
      CREATE PROPERTY metadata -> std::str;
      CREATE REQUIRED PROPERTY name -> std::str;
  };
  CREATE TYPE default::Contract {
      CREATE MULTI LINK assets -> default::Asset;
      CREATE REQUIRED PROPERTY address -> std::str;
  };
  CREATE TYPE default::Type {
      CREATE MULTI LINK contracts -> default::Contract;
      CREATE REQUIRED PROPERTY name -> std::str;
  };
  CREATE TYPE default::Subscription {
      CREATE REQUIRED MULTI LINK types -> default::Type;
      CREATE REQUIRED PROPERTY blockchain -> std::str;
      CREATE REQUIRED PROPERTY lastMint -> std::str;
      CREATE REQUIRED PROPERTY name -> std::str;
  };
  CREATE TYPE default::User {
      CREATE REQUIRED PROPERTY email -> std::str;
      CREATE REQUIRED PROPERTY ethAddress -> std::str;
      CREATE PROPERTY foAddress -> std::str;
      CREATE REQUIRED PROPERTY userName -> std::str;
      CREATE PROPERTY xcpAddress -> std::str;
  };
};
