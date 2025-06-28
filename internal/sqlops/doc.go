// Package sqlops is a collection of operations related to the SQL operations.
//
// Database Type Differences:
//
// Understanding the differences in SQL concepts and data types across
// various database systems (SQLite, PostgreSQL, MySQL) is crucial for
// writing portable and efficient SQL operations. Below is a summary:
//
//		Concept                       SQLite           PostgreSQL                     MySQL                           Notes
//		---------------------------   -------------    -------------------------      -------------------------       ----------------------------------------------------------------------
//		Integer                       INTEGER          INTEGER, INT, BIGINT,          INT, BIGINT, TINYINT,           SQLite's INTEGER is quite flexible and can store various integer sizes.
//		                                               SMALLINT                       SMALLINT, MEDIUMINT             PostgreSQL and MySQL offer more specific integer types for different ranges.
//
//		Text/String                   TEXT             VARCHAR(n), TEXT,              VARCHAR(n), TEXT,               TEXT in SQLite is typically variable-length. In PostgreSQL and MySQL,
//		                                               CHAR(n)                        CHAR(n)                         TEXT is for very long strings, while VARCHAR(n) is for variable-length
//		                                                                                                              strings up to n characters. CHAR(n) is fixed-length.
//
//		Numbers (Decimal/Floating)    REAL, NUMERIC    NUMERIC(p,s), DECIMAL(p,s),    DECIMAL(p,s), NUMERIC(p,s),     REAL in SQLite is a floating-point number. NUMERIC(p,s)/DECIMAL(p,s) are for
//		                                               REAL, DOUBLE PRECISION         FLOAT, DOUBLE                   exact precision (p=precision, s=scale) and are widely supported.
//		                                                                                                              FLOAT and DOUBLE are for approximate floating-point numbers.
//
//		Boolean                       INTEGER          BOOLEAN, BOOL                  TINYINT(1)                      SQLite doesn't have a native boolean type, often using INTEGER instead.
//		                              (0 for false,                                   (0 for false,                   MySQL often uses TINYINT(1) for boolean, and PostgreSQL has a dedicated BOOLEAN type.
//	                                1 for true)                                     1 for true)
//
//
//		Date/Time                     TEXT, INTEGER,   DATE, TIME, TIMESTAMP,         DATE, TIME, DATETIME,           SQLite stores dates/times as text (ISO8601 strings), integers (Unix epoch time),
//		                              REAL             TIMESTAMPTZ                    TIMESTAMP                       or real numbers (Julian day numbers). PostgreSQL and MySQL have dedicated and more
//		                                                                                                              robust date/time types, including options for time zones (TIMESTAMPTZ in PostgreSQL).
//
//		Binary Data                   BLOB             BYTEA                          BLOB, TINYBLOB,                 All support binary large objects.
//		                                                                              MEDIUMBLOB, LONGBLOB
//
// # Statemebt oarameter binding
//
// Database     Positional Anonymous    Positional Numbered    Named (Native SQL).     Named (Client/Driver specific)
// ----------   ----------------------  --------------------   ---------------------   -------------------------------
// SQLite.      ?                       ?N                     :name, @name, $name.   Yes (often supports all)
// MySQL        ?                       No                     No                     Yes (common in client libraries)
// PostgreSQL   No                      $N                     No                     Yes (common in client libraries)
//
// Note: This table provides a general overview. Specific driver
// implementations or database versions might have nuances. Always
// consult the official documentation for precise details.
//
// For more detailed information on SQL operations, refer to the
// functions and types defined within this package.
package sqlops
