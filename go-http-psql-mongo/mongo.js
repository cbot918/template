const CryptoJS = require("crypto-js");
const mongoose = require("mongoose");

const user = "yale";
const pw = "12345";
const db = "testdb";
const port = "27019";

const MONGO_URL = `mongodb://${user}:${pw}@localhost:${port}/${db}?retryWrites=true&w=majority`;
const PASS_SEC = "abcde12345";
(async () => {
  mongoose
    .connect(MONGO_URL)
    .then(() => console.log("DB Connection Successfull!"))
    .catch((err) => {
      console.log(err);
    });

  const UserSchema = new mongoose.Schema(
    {
      username: { type: String, required: true, unique: true },
      email: { type: String, required: true, unique: true },
      password: { type: String, required: true },
      isAdmin: {
        type: Boolean,
        default: false,
      },
      img: { type: String },
    },
    { timestamps: true }
  );

  User = mongoose.model("User", UserSchema);

  const newUser = new User({
    username: "yale918",
    email: "yale918@gmail.com",
    password: CryptoJS.AES.encrypt("12345", PASS_SEC).toString(),
  });

  try {
    const savedUser = await newUser.save();
    console.log(" user saved ...");
  } catch (err) {
    console.log(" user saved error ...");
    console.log(err);
  }
})();
