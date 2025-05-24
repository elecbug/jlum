using System.Text.Json.Serialization;

namespace JLUM.Data
{
    public struct Information
    {
        public string JournalName { get; set; }
        public int Volume { get; set; }
        public int Number { get; set; }
        public string Issue { get; set; }
        public Page Page { get; set; }
        public JDate Date { get; set; }
        public Identifier Identifiers { get; set; }
        public DB DB { get; set; }

        public static Information Empty => new Information
        {
            JournalName = "",
            Volume = -1,
            Number = -1,
            Issue = "",
            Page = Page.Unknown,
            Date = JDate.Unknown,
            Identifiers = Identifier.Empty,
            DB = DB.None,
        };

        public override string ToString()
        {
            string result = "";
            List<string> results = [];

            if (!string.IsNullOrEmpty(JournalName))
                results.Add($"{JournalName}");
            if (Volume != -1)
                results.Add($"Vol. {Volume}");
            if (Number != -1)
                results.Add($"No. {Number}");
            if (!string.IsNullOrEmpty(Issue))
                results.Add($"Issue {Issue}");
            if (Page.ToString() != "")
                results.Add(Page.ToString());
            if (Date.ToString() != "")
                results.Add(Date.ToString());
            if (Identifiers.ToString() != "")
                results.Add(Identifiers.ToString());

            result = string.Join(", ", results);

            if (DB != DB.None)
            {
                List<string> dbResults = [];

                if (DB.HasFlag(DB.KCI))
                    dbResults.Add("KCI");
                if (DB.HasFlag(DB.SCOPUS))
                    dbResults.Add("SCOPUS");
                if (DB.HasFlag(DB.SCI))
                    dbResults.Add("SCI");
                if (DB.HasFlag(DB.SCIE))
                    dbResults.Add("SCIE");
                if (DB.HasFlag(DB.IEEE))
                    dbResults.Add("IEEE");

                result += $"({string.Join(", ", dbResults)})";
            }

            return result;
        }
    }

    public struct Identifier
    {
        public string DOI { get; set; }
        public string ISBN { get; set; }
        public string ISSN { get; set; }
        public string EISSN { get; set; }

        public static Identifier Empty => new Identifier
        {
            DOI = "",
            ISBN = "",
            ISSN = "",
            EISSN = ""
        };

        public override string ToString()
        {
            List<string> results = [];
            if (!string.IsNullOrEmpty(DOI))
                results.Add($"DOI: {DOI}");
            if (!string.IsNullOrEmpty(ISBN))
                results.Add($"ISBN: {ISBN}");
            if (!string.IsNullOrEmpty(ISSN))
                results.Add($"ISSN: {ISSN}");
            if (!string.IsNullOrEmpty(EISSN))
                results.Add($"EISSN: {EISSN}");

            return string.Join(", ", results);
        }
    }

    public struct Page
    {
        public int Start { get; set; }
        public int End { get; set; }

        [JsonIgnore]
        public int Length => End - Start + 1;

        public static Page Unknown => new Page
        {
            Start = -1,
            End = -1,
        };

        public override string ToString()
        {
            if (Start == -1 && End == -1)
                return "";
            else if (Length == 1)
                return $"pp. {Start.ToString()}";
            else
                return $"pp. {Start}-{End}";
        }
    }

    public struct JDate
    {
        private static readonly string[] MonthNames = new[]
        {
            "January", "February", "March", "April", "May", "June",
            "July", "August", "September", "October", "November", "December"
        };

        public int Year { get; set; }
        public int Month { get; set; }

        [JsonIgnore]
        public int MonthIndex => Month - 1;
        [JsonIgnore]
        public string MonthFullName => MonthNames[MonthIndex];
        [JsonIgnore]
        public string MonthShortName => MonthNames[MonthIndex].Substring(0, 3) + ".";

        public static JDate Unknown => new JDate
        {
            Year = -1,
            Month = -1,
        };

        public override string ToString()
        {
            if (Year == -1 && Month == -1)
                return "";
            else if (Month == -1)
                return Year.ToString();
            else
                return $"{MonthShortName} {Year}";
        }
    }

    public enum DB
    {
        None = 0,
        KCI = 1,
        SCOPUS = 2,
        SCI = 4,
        SCIE = 8,
        IEEE = 16,
    }
}