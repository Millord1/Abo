CREATE MIGRATION m1mciecc3zry3v7fwaitqvadvemufxwgs4xjxa4wpk2ahkfh3x6uqa
    ONTO m1z5fccddnrzv4w34gcfut6hmjiv2vz775g3xuwjicw6lb7ij27z2q
{
  ALTER TYPE default::Asset {
      ALTER PROPERTY assetId {
          CREATE CONSTRAINT std::exclusive;
      };
  };
  ALTER TYPE default::Collection {
      ALTER PROPERTY collectionId {
          CREATE CONSTRAINT std::exclusive;
      };
  };
  ALTER TYPE default::Contract {
      ALTER PROPERTY address {
          CREATE CONSTRAINT std::exclusive;
      };
  };
  ALTER TYPE default::MintLog {
      ALTER PROPERTY txHash {
          CREATE CONSTRAINT std::exclusive;
      };
  };
  ALTER TYPE default::Subscription {
      ALTER PROPERTY name {
          CREATE CONSTRAINT std::exclusive;
      };
  };
  ALTER TYPE default::Type {
      ALTER PROPERTY name {
          CREATE CONSTRAINT std::exclusive;
      };
  };
  ALTER TYPE default::User {
      ALTER PROPERTY email {
          CREATE CONSTRAINT std::exclusive;
      };
  };
};
