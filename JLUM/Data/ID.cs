using System.Security.Cryptography;

namespace JLUM.Data
{
    public class ID
    {
        public enum Type
        {
            Unknown = 0,
            UID = 1,
            PID = 2,
            AID = 3,
        }

        private byte[] _id;
        private Type _type;

        private ID(Type idType)
        {
            _id = new byte[32];
            _type = idType;

            RandomNumberGenerator.Fill(_id);
        }
        public static ID GenerateID(Type idType)
        {
            return new ID(idType);
        }

        public override string ToString()
        {
            return $"{TypeToString()}{Convert.ToHexString(_id)}";
        }
        public static ID FromString(string idString)
        {
            if (idString.Length != 66)
                throw new ArgumentException("ID string must be 66 characters long.");
            Type idType = idString[0] switch
            {
                'U' => Type.UID,
                'P' => Type.PID,
                'A' => Type.AID,
                _ => Type.Unknown
            };
            byte[] idBytes = Convert.FromHexString(idString.Substring(1, 64));
            ID id = new ID(idType);
            id._id = idBytes;
            return id;
        }

        private string TypeToString()
        {
            return _type switch
            {
                Type.UID => "U",
                Type.PID => "P",
                Type.AID => "A",
                _ => "Unknown"
            };
        }
    }
}
