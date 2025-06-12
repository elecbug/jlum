namespace JLUM.Data
{
    public struct Author
    {
        public string Name { get; set; }
        public string Affiliation { get; set; }
        public string Email { get; set; }
        public AuthorRole AuthorRole { get; set; }

        public static Author Empty => new Author
        {
            Name = "",
            Affiliation = "",
            Email = "",
            AuthorRole = AuthorRole.Unknown
        };

        public string ToIEEE()
        {
            return $"{Name}";
        }
    }

    public enum AuthorRole
    {
        Unknown = 0,
        FirstAuthor = 1,
        SecondAuthor = 2,
        CoAuthor = 3,
        CorrespondingAuthor = 4,
    }
}