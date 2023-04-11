CREATE MIGRATION m1z5fccddnrzv4w34gcfut6hmjiv2vz775g3xuwjicw6lb7ij27z2q
    ONTO m14rruqf2oruky3gs2ygcqljhipf5fdvrji6syior6upfe5pwmkjea
{
  CREATE TYPE default::Collection {
      CREATE REQUIRED PROPERTY collectionId -> std::str;
      CREATE REQUIRED PROPERTY description -> std::str;
      CREATE REQUIRED PROPERTY imgUrl -> std::str;
      CREATE REQUIRED PROPERTY name -> std::str;
  };
  ALTER TYPE default::Asset {
      CREATE LINK collection -> default::Collection;
      CREATE MULTI LINK contracts -> default::Contract;
  };
  ALTER TYPE default::Contract {
      DROP LINK assets;
  };
  ALTER TYPE default::Contract {
      CREATE REQUIRED PROPERTY blockchain -> std::str {
          SET REQUIRED USING ('klaytn');
      };
  };
  ALTER TYPE default::Contract {
      CREATE REQUIRED PROPERTY standard -> std::str {
          SET REQUIRED USING ('ERC721');
      };
  };
  CREATE TYPE default::MintLog {
      CREATE REQUIRED LINK asset -> default::Asset;
      CREATE REQUIRED LINK contract -> default::Contract;
      CREATE REQUIRED PROPERTY date -> std::str;
      CREATE REQUIRED PROPERTY txHash -> std::str;
  };
  ALTER TYPE default::User {
      CREATE MULTI LINK mintLogs -> default::MintLog;
      CREATE LINK hasType -> default::Type;
  };
  ALTER TYPE default::Type {
      CREATE MULTI LINK assets -> default::Asset {
          CREATE PROPERTY quantity -> std::int32;
      };
  };
  ALTER TYPE default::Type {
      DROP LINK contracts;
  };
};
