// Create user
dbAdmin = db.getSiblingDB("admin");
dbAdmin.createUser({
    user: "authUser",
    pwd: "authPassword",
    roles: [{ role: "userAdminAnyDatabase", db: "admin" }],
    mechanisms: ["SCRAM-SHA-1"],
});

// Authenticate user
dbAdmin.auth({
    user: "authUser",
    pwd: "authPassword",
    mechanisms: ["SCRAM-SHA-1"],
    digestPassword: true,
});

// Create DB and collection
db = new Mongo().getDB("auth");
db.createCollection("users", { capped: false });

// Create index for google id
db.getCollection('users').createIndex({ googleId: 'text' }, { name: 'googleId', unique: true });