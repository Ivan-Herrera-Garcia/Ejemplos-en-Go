USE [ESTUDIANTES] /*CREAR BASE DE DATOS ESTUDIANTES*/
GO

/****** Object:  Table [dbo].[INFORMACION]    Script Date: 19/01/2023 10:28:30 p. m. ******/
SET ANSI_NULLS ON
GO

SET QUOTED_IDENTIFIER ON
GO

CREATE TABLE [dbo].[INFORMACION](
	[NUM_CONTROL] [char](8) NOT NULL,
	[APE_P] [varchar](50) NOT NULL,
	[APE_M] [varchar](50) NOT NULL,
	[NOMBRE] [varchar](50) NOT NULL,
	[EDAD] [smallint] NULL,
	[GENERO] [varchar](50) NULL,
 CONSTRAINT [PK_INFORMACION_NUM_CONTROL] PRIMARY KEY CLUSTERED 
(
	[NUM_CONTROL] ASC
)WITH (PAD_INDEX = OFF, STATISTICS_NORECOMPUTE = OFF, IGNORE_DUP_KEY = OFF, ALLOW_ROW_LOCKS = ON, ALLOW_PAGE_LOCKS = ON, OPTIMIZE_FOR_SEQUENTIAL_KEY = OFF) ON [PRIMARY]
) ON [PRIMARY]
GO

INSERT INTO INFORMACION VALUES ('19134587','Ochoa','Rodriguez','Juan',25,'Male')
INSERT INTO INFORMACION VALUES ('19134574','Garcia','Hernandez','Johan',33,'Male')
INSERT INTO INFORMACION VALUES ('19134555','Castillo','Coronado','Marcelo',16,'Male')
GO
