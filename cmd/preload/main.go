package main

import (
	"fmt"

	"github.com/Orfeo42/admin-panel/data"
	"github.com/Orfeo42/admin-panel/db"
)

func main() {
	updateSchema()
	inizializeCustomersData()
}

func updateSchema() error {
	fmt.Println("Updating Schema...")
	dbInstance, errConnection := db.GetInstance()

	if errConnection != nil {
		return errConnection
	}
	err := dbInstance.AutoMigrate(
		&data.Customer{},
		&data.Invoice{},
		&data.InvoiceRow{},
		&data.Order{},
		&data.OrderRow{},
	)
	if err != nil {
		fmt.Println("Error in Updating Schema")
		return err
	}
	fmt.Println("Schema Updated!")
	return nil
}

func inizializeCustomersData() ([]data.Customer, error) {
	customerList := []data.Customer{
		{
			Name: "03 PALSTIC  S.R.L.",
		},
		{
			Name: "2.G. ORO DI BALDONI GIANFRANCO",
		},
		{
			Name: "2009 S.A.S. DI BOBE MARIANA",
		},
		{
			Name: "2P ENGINEERING SRL",
		},
		{
			Name: "3 C. DI CAPACCI MASSIMILIANO",
		},
		{
			Name: "3.F. AGRICOLA S.R.L. SOCIETA AGRICOLA",
		},
		{
			Name: "4 STARS EVENTI DI FREQUENTINI",
		},
		{
			Name: "4C AREZZO S.N.C.",
		},
		{
			Name: "A.C.D. SAN GIOVANNI-AREZZO",
		},
		{
			Name: "A.L. S.R.L.",
		},
		{
			Name: "A.L. SAS DI NEDA A.F. & C.",
		},
		{
			Name: "A.M.IMPIANTI S.R.L",
		},
		{
			Name: "A.S.D CORTONA CAMUCIA",
		},
		{
			Name: "A.S.D RONDA GHIBELLINA TEAM",
		},
		{
			Name: "A.S.D. C.S. VITIANO",
		},
		{
			Name: "A.S.D. CAMUCIA CORTONA",
		},
		{
			Name: "A.S.D. CORTONA BOCCE",
		},
		{
			Name: "A.S.D. G.S. LA ROCCA",
		},
		{
			Name: "A.S.D. GRUPPO SP.SAN LORENZO",
		},
		{
			Name: "A.S.D. LA CHIANINA",
		},
		{
			Name: "A.S.D. MONSIGLIOLO",
		},
		{
			Name: "A.S.D. SOLENGO MONTANARE",
		},
		{
			Name: "A.S.D.\"EMOZIONI\"A.C.S.I. CLUB",
		},
		{
			Name: "A.S.D.FRATTA S.CATERINA",
		},
		{
			Name: "A7 ALIMENTARI DI SUTURI E",
		},
		{
			Name: "AC HOTEL AREZZO S.R.L.",
		},
		{
			Name: "ACCADEMIA DEGLI ARDITI",
		},
		{
			Name: "ACETIFICIO ARETINO SRL",
		},
		{
			Name: "ACQUA & CO. DI FIGLIUZZI",
		},
		{
			Name: "ACQUA B. LIBERATORI.A.&.C",
		},
		{
			Name: "ACQUA SAN BENEDETTO S.P.A.",
		},
		{
			Name: "AD S.R.L.",
		},
		{
			Name: "ADAT CONSORZIO DISTRIBUTORI",
		},
		{
			Name: "AEMME S.N.C. DI MENCI MAURIZIO",
		},
		{
			Name: "AEREO CLUB SERRISTORI",
		},
		{
			Name: "AERRE SOCETA  A RESPONSABI",
		},
		{
			Name: "AF DI FRATINI ALESSANDRO",
		},
		{
			Name: "AF. DI FRATTINI ALESANDRO",
		},
		{
			Name: "AF.UT DI BIAGIOTTI GIANFRANCO",
		},
		{
			Name: "AGIP 4645 LUCIGNANO EST",
		},
		{
			Name: "AGNELLI AGOSTINO",
		},
		{
			Name: "AGNELLI SIMONE BAR 39",
		},
		{
			Name: "AGRICOLA VTI DI LUCA",
		},
		{
			Name: "AGRIO SNC DI MANITIO SUNSHHINE",
		},
		{
			Name: "AGRITURISMO I PAGLIAI DI MAURA",
		},
		{
			Name: "AGRITURISMO LA MAESTA DI",
		},
		{
			Name: "AGRITURISMO VALLE DAME DI",
		},
		{
			Name: "AGROINVEST S.R.L.",
		},
		{
			Name: "AHMAD HASNAT NEW KEBAB",
		},
		{
			Name: "AHMAD WAQAS A.D. KEBAB E PIZZA",
		},
		{
			Name: "AHMED NEW K.F.C. KEEBAB",
		},
		{
			Name: "AHMED.S KEBAB P.DEL SOTTILE SC",
		},
		{
			Name: "AIGRENIS S.R.L",
		},
		{
			Name: "AION CULTURA",
		},
		{
			Name: "AION CULTURA PICCOLA SOCIETA",
		},
		{
			Name: "AL FORNO DE PIAZZA DI IMPARATO",
		},
		{
			Name: "ALBERGO OSTERIA TOTO",
		},
		{
			Name: "ALBERGO RISTORANTE \"SAURO\" DI",
		},
		{
			Name: "ALBERGO RISTORANTE FRATELLI",
		},
		{
			Name: "ALBERGO RISTORANTE PORTOLE",
		},
		{
			Name: "ALBERTI BRYAN CAFFE DEL MUSEO",
		},
		{
			Name: "ALEMAS S.R.L.",
		},
		{
			Name: "ALESSANDRO BATTOLLA",
		},
		{
			Name: "ALI DANISH KEKAB",
		},
		{
			Name: "ALI DANISH KEKAB 2",
		},
		{
			Name: "ALI KEBAB & PIZZA DI ALI ZULFI",
		},
		{
			Name: "ALI KEBAB DI MURTAZA GHULAM",
		},
		{
			Name: "ALIBABA KEBAB AND GRILL DI",
		},
		{
			Name: "ALIBY CAFE' S.N.C. DI BANINI L",
		},
		{
			Name: "ALIDA COIFFEUR DI BRUSCHI",
		},
		{
			Name: "ALIME DA NONNA ANNA DI MALZONE",
		},
		{
			Name: "ALIME. DI TUFO JENINA & C SAS",
		},
		{
			Name: "ALIMENTARI BANELLI PIETRO",
		},
		{
			Name: "ALIMENTARI BAR \"IL CHIANINO\"",
		},
		{
			Name: "ALIMENTARI CATORCIONI VERA",
		},
		{
			Name: "ALIMENTARI CUSERI DONATO",
		},
		{
			Name: "ALIMENTARI DI GARZI DANIELA",
		},
		{
			Name: "ALIMENTARI EREDI SEGANTINI",
		},
		{
			Name: "ALIMENTARI FRATELLI BIAGIANTI",
		},
		{
			Name: "ALIMENTARI GALOPPI SERGIO",
		},
		{
			Name: "ALIMENTARI IL NEGOZIETTO DI",
		},
		{
			Name: "ALIMENTARI LE 3 ESSE SNC DI",
		},
		{
			Name: "ALIMENTARI M.MARKET PIETRAIA",
		},
		{
			Name: "ALIMENTARI MAJA DI CAMERINI",
		},
		{
			Name: "ALIMENTARI MANERCHIA",
		},
		{
			Name: "ALIMENTARI MASCAGNI S.A.S.",
		},
		{
			Name: "ALIMENTARI MENCI MARIA GRAZIA",
		},
		{
			Name: "ALIMENTARI PRESENTINI PAOLO",
		},
		{
			Name: "ALIMENTARI SALVADORI S.A.S DI",
		},
		{
			Name: "ALIMENTARI SIMONA DI RAGNI",
		},
		{
			Name: "ALIMENTARI TIEZZI FRANCA",
		},
		{
			Name: "ALIMENTARI ZUCCHINI SAVINA",
		},
		{
			Name: "ALISHAN-KEBAB CUCINA INDIANA",
		},
		{
			Name: "ALLIANCE MEDICAL S.R.L.",
		},
		{
			Name: "ALLIED DOMECQ ITALIA S.P.A.",
		},
		{
			Name: "ALMAS COSTRUZIONI SRL",
		},
		{
			Name: "ALTERNATIVA SERVICE DI CECCARE",
		},
		{
			Name: "AMARANTE CARBURANTE E SERVIZI DI S. AMARANTE",
		},
		{
			Name: "AMC S.R.L.",
		},
		{
			Name: "AMC SRL BAR",
		},
		{
			Name: "AMICI DEL RE DELLA MACCHIA",
		},
		{
			Name: "AMICI FREE S.R.L.",
		},
		{
			Name: "AMM-CARS DI RISPOLI UMBERTO",
		},
		{
			Name: "AMMINISTRAZIONE PROVINCIALE",
		},
		{
			Name: "AN.DE.P DI ANTONIO DE CECCO",
		},
		{
			Name: "ANDROMEDA DI BERNI GIAN LUCA",
		},
		{
			Name: "ANDROMEDA S.R.L.",
		},
		{
			Name: "ANDRY BAR S.N.C. DI",
		},
		{
			Name: "ANGOLO DEL PANE DI HOXRA",
		},
		{
			Name: "ANGORI DOTT.GIAN LUCA",
		},
		{
			Name: "ANNIBAL S.N.C. DI STRAPAGHETTI",
		},
		{
			Name: "ANTE LABOREM\" S.C.",
		},
		{
			Name: "ANTICA CONCERIA S.R.L.",
		},
		{
			Name: "ANTICA LAVORAZIONE CARNI DI",
		},
		{
			Name: "ANTICA TRATTORIA DAL 1904",
		},
		{
			Name: "ANTICA TRATTORIA LA FOCE DI",
		},
		{
			Name: "ANTICHI SAPORI ALIMENTARI",
		},
		{
			Name: "ANTICO CAFFE' LA POSTA S.R.L.",
		},
		{
			Name: "ANTINORI SOC.AGRICOLA A RL",
		},
		{
			Name: "ANTONELLI S.R.L.",
		},
		{
			Name: "ANTONIO CARACUTA IL GIRONE DEI",
		},
		{
			Name: "APICE S.R.L.",
		},
		{
			Name: "AR.BERE DI PIANTINI PAOLO",
		},
		{
			Name: "ARCA ENEL ASSOCIAZIONE NAZIONE",
		},
		{
			Name: "ARCANGELA L'ALIMENTARE",
		},
		{
			Name: "ARCHITEC S.R.L.",
		},
		{
			Name: "ARCI AREZZO",
		},
		{
			Name: "AREZZO WAVE MANAGEMENT",
		},
		{
			Name: "ARLECCHINO ENTE DEL TERSO SETTORE",
		},
		{
			Name: "ARMEC S.N.C. DI MEACCI L.",
		},
		{
			Name: "ARMERIA DEL ZONZO DI CAMORRI F",
		},
		{
			Name: "ARNOLFO DRINK S.N.C.",
		},
		{
			Name: "ARREDAMENTI ARTIGIANI S.N.C",
		},
		{
			Name: "ARREDO DI PIETRA S.R.L.",
		},
		{
			Name: "ARRIGUCCI MILLI",
		},
		{
			Name: "ARTE & RICAMO S.R.L.",
		},
		{
			Name: "ARTE BIANCA CORTONESE S.N.C.",
		},
		{
			Name: "ARTE DEL GELATO",
		},
		{
			Name: "ARTE DOLCE PASTICCERIA",
		},
		{
			Name: "ASS. CULTURALE \"IL CARRO\"",
		},
		{
			Name: "ASS. PROLOCO VAL DEL BAGNORO",
		},
		{
			Name: "ASS. TERRAZZA DEL FOLK",
		},
		{
			Name: "ASS.SPORTIVA IL TOPPALE",
		},
		{
			Name: "ASSICLANIS S.R.L.",
		},
		{
			Name: "ASSOC. LA VALLA IN FESTA",
		},
		{
			Name: "ASSOCI.CITTADINI DI BORGHETTO",
		},
		{
			Name: "ASSOCIAZION IL CORPO NEL MONDO",
		},
		{
			Name: "ASSOCIAZIONE \"AMICI DI VADA\"",
		},
		{
			Name: "ASSOCIAZIONE <<IL PALAZZO>>",
		},
		{
			Name: "ASSOCIAZIONE AMICI DI SIMONE",
		},
		{
			Name: "ASSOCIAZIONE CANTIERE RUSTICI",
		},
		{
			Name: "ASSOCIAZIONE CICLISTICA",
		},
		{
			Name: "ASSOCIAZIONE COMETA",
		},
		{
			Name: "ASSOCIAZIONE COMITATO FESTEGGI",
		},
		{
			Name: "ASSOCIAZIONE CULTUR. ONTHEMOVE",
		},
		{
			Name: "ASSOCIAZIONE CULTURALE ARTISTI",
		},
		{
			Name: "ASSOCIAZIONE CULTURALE CAPOTRA  VEKILOWATT",
		},
		{
			Name: "ASSOCIAZIONE CULTURALE DEL",
		},
		{
			Name: "ASSOCIAZIONE CULTURALE IL",
		},
		{
			Name: "ASSOCIAZIONE CULTURALE SAGRA",
		},
		{
			Name: "ASSOCIAZIONE CULTURALE VAL DI",
		},
		{
			Name: "ASSOCIAZIONE DIVINA CLUB",
		},
		{
			Name: "ASSOCIAZIONE ETRURIA ANIMALS",
		},
		{
			Name: "ASSOCIAZIONE GLIO.MA",
		},
		{
			Name: "ASSOCIAZIONE LA NUIT",
		},
		{
			Name: "ASSOCIAZIONE ONLUS AMICI DI",
		},
		{
			Name: "ASSOCIAZIONE PRO LOCO CENTOIA",
		},
		{
			Name: "ASSOCIAZIONE PRO-LOCO BETTOLLE",
		},
		{
			Name: "ASSOCIAZIONE PRO-LOCO RIGOMAGN",
		},
		{
			Name: "ASSOCIAZIONE PRO-LOCO TUORO",
		},
		{
			Name: "ASSOCIAZIONE PRO-LOCO TURISTIC",
		},
		{
			Name: "ASSOCIAZIONE RIONE MALPASSO",
		},
		{
			Name: "ASSOCIAZIONE RIONE OLIVETO",
		},
		{
			Name: "ASSOCIAZIONE SPORTIVA DILETTANTISTICA MAESTA",
		},
		{
			Name: "ASSOCIAZIONE SPORTIVA RICREATI",
		},
		{
			Name: "ATHENA SOCIETA COOPERATIVA",
		},
		{
			Name: "ATLETICO VALDICHIANA DILETTANT",
		},
		{
			Name: "ATTILIO GORI S.R.L.",
		},
		{
			Name: "AUTOCARROZZERIA VALENTINI",
		},
		{
			Name: "AUTOCORIS s.r.l",
		},
		{
			Name: "AUTOFFICINA 2000 BORZI PAOLO",
		},
		{
			Name: "AUTOFFICINA QUERCE S.R.L.",
		},
		{
			Name: "AUTOGRILL PARIS SNC TRATTORIA",
		},
		{
			Name: "AUTOGRILL S.P.A.",
		},
		{
			Name: "AUTOSCUOLA  VESTRINI",
		},
		{
			Name: "AUTOSERVIZI VOLPI S.R.L.",
		},
		{
			Name: "AUTOTRASPORTI LEPRI ANDREA",
		},
		{
			Name: "AUTOWEB SRL GIANLUCA CALVANI",
		},
		{
			Name: "AVIGNONESI S.R.L. SOCIETA",
		},
		{
			Name: "AVVISATI FRANCESCO",
		},
		{
			Name: "AZ AGRITURISTICA IL BOSCO",
		},
		{
			Name: "AZ. AGR. MERAVIGLI PATRIZZIA",
		},
		{
			Name: "AZ. AGRICOLA CHIOVOLONI",
		},
		{
			Name: "AZ.AGR.MERAVIGLI PATRIZIA",
		},
		{
			Name: "AZ.AGR.MORETTI SOC AGR.S.",
		},
		{
			Name: "AZ.AGRICOLA SANTO STEFANO DI",
		},
		{
			Name: "AZIE. AGRICOLA DI GENCA AGIOLO",
		},
		{
			Name: "AZIENDA  AGRICOLA   MARCO MONTINI",
		},
		{
			Name: "AZIENDA AGRARIA ROSELLI E. e",
		},
		{
			Name: "AZIENDA AGRICOLA BALDESCHI MATTEO",
		},
		{
			Name: "AZIENDA AGRICOLA BALDETTI",
		},
		{
			Name: "AZIENDA AGRICOLA BEMOCCOLI",
		},
		{
			Name: "AZIENDA AGRICOLA BROCCHI",
		},
		{
			Name: "AZIENDA AGRICOLA FRA.BRIGANTI",
		},
		{
			Name: "AZIENDA AGRICOLA GIROMAGI",
		},
		{
			Name: "AZIENDA AGRICOLA MENCACCI",
		},
		{
			Name: "AZIENDA AGRICOLA MORETTI",
		},
		{
			Name: "AZIENDA AGRICOLA POGGIO AL",
		},
		{
			Name: "AZIENDA AGRICOLA TENUTA",
		},
		{
			Name: "AZIENDA AGRICOLA VILLA LA",
		},
		{
			Name: "AZIENDA AGRICOLA VITI EMANUELE",
		},
		{
			Name: "AZIENDA AGRICOLA VITI UMBERTO",
		},
		{
			Name: "AZZIENDA AGRICOLA SANTINI",
		},
		{
			Name: "B & B CASA CAPANNI",
		},
		{
			Name: "B & P AUTOMOTIVE S.R.L.",
		},
		{
			Name: "B J RACING TEAM",
		},
		{
			Name: "B.F.AGRICOLA S.R.L.SOCIETA",
		},
		{
			Name: "B.G. DI BONDI GIACOMO",
		},
		{
			Name: "B.M. DI BARTOLOZZI MARIA",
		},
		{
			Name: "B.R.G. S.R.L.",
		},
		{
			Name: "B.T.TRUCKTRANS S.R.L.",
		},
		{
			Name: "BACCHETTA ANNA LISA",
		},
		{
			Name: "BACCI MOTORS S.R.L.",
		},
		{
			Name: "BADIA DI POMAIO S.R.L.",
		},
		{
			Name: "BADIACCIA S.A.S. DI PIOMBINI",
		},
		{
			Name: "BADII SILVIA",
		},
		{
			Name: "BALDETTI GIORGIO S.R.L.",
		},
		{
			Name: "BALDI CARLO BLUE'S BAR",
		},
		{
			Name: "BALDOLUNGHI MASSIMO",
		},
		{
			Name: "BANCA ETRURIA",
		},
		{
			Name: "BANCA VALDICHIANA CREDITO COOP",
		},
		{
			Name: "BANCHELLI ELIAS",
		},
		{
			Name: "BANFI S.R.L.",
		},
		{
			Name: "BANGLA ALIMENTARI SNC",
		},
		{
			Name: "BAR 1990 DI NICOLAE LUMINITA",
		},
		{
			Name: "BAR 2000 DI BOTTI DEBORA E",
		},
		{
			Name: "BAR 500 S.R.L.",
		},
		{
			Name: "BAR ACQUILA S.A.S. DI CENAJ",
		},
		{
			Name: "BAR AL CAFFE' DI VALENTE",
		},
		{
			Name: "BAR ALIMENTARI BELMAR S.N.C.",
		},
		{
			Name: "BAR AMARANTO POINT DI TESTI",
		},
		{
			Name: "BAR ANGELLS LOUNGE DI",
		},
		{
			Name: "BAR ANTIPORTO DI LOBINA ELISA",
		},
		{
			Name: "BAR BELLUCCI ANNA BAR OSPEDALE",
		},
		{
			Name: "BAR CAFFE GALAXIA DI",
		},
		{
			Name: "BAR CAFFE LE TORRI",
		},
		{
			Name: "BAR CAFFE SIGNORELLI S.N.C.",
		},
		{
			Name: "BAR CAFFE'PIT STOP DI PARIGI",
		},
		{
			Name: "BAR CENTRALE D. SPORT RIGUTINO",
		},
		{
			Name: "BAR CENTRALE DI GJETA ELVIRA",
		},
		{
			Name: "BAR CIN CIN S.R.L",
		},
		{
			Name: "BAR CONTROVENTO S.N.C.DI",
		},
		{
			Name: "BAR DEI FILARMONICI DI GHEZZI&",
		},
		{
			Name: "BAR DEL PASSAGGIO DI NICOLETTI SERENA",
		},
		{
			Name: "BAR DEL RICCIO S.R.L.",
		},
		{
			Name: "BAR DELL INDICATORE S.R.L",
		},
		{
			Name: "BAR DELLE MURA S.N.C. DI",
		},
		{
			Name: "BAR DELLO SPORT DI BRINI R.",
		},
		{
			Name: "BAR DOLCE CAFFE MIHAI",
		},
		{
			Name: "BAR DOLCE SALATO DI ANTONELLA",
		},
		{
			Name: "BAR ELISA DI BRANCARDI",
		},
		{
			Name: "BAR ELISIR VALERIA DI MAMELI",
		},
		{
			Name: "BAR EUROPA DI PAOLA MOSCARDO",
		},
		{
			Name: "BAR EXEDO DI CIAMPA LUIGI & C.",
		},
		{
			Name: "BAR FASHION CAFFE DI HARABALA",
		},
		{
			Name: "BAR FENICE DI ARMEANU IULIANA",
		},
		{
			Name: "BAR G.O. DI GROMOVA OLGA",
		},
		{
			Name: "BAR GARIBALDI DI ZENNARO RENZO",
		},
		{
			Name: "BAR GELATERIA MIRAGGIO DI",
		},
		{
			Name: "BAR GELATERIA NUY BOYS DI",
		},
		{
			Name: "BAR GREEDY ARK GROUP SALS",
		},
		{
			Name: "BAR I GIRASOLI DI BRUSCHI",
		},
		{
			Name: "BAR IL BARRINO S.N.C. DI",
		},
		{
			Name: "BAR IL MOLO S.N.C DI DIMITROV",
		},
		{
			Name: "BAR ITALIA DI ALI SHAHID",
		},
		{
			Name: "BAR L'INCONTRO S.R.L.",
		},
		{
			Name: "BAR LA FORTEZZA DI BERNARDINI",
		},
		{
			Name: "BAR LA MAESTA S.A.S",
		},
		{
			Name: "BAR LA PERLA",
		},
		{
			Name: "BAR LA PIAZZA DI STRAPAGHETTI",
		},
		{
			Name: "BAR LA PIAZZETTA DE ORI DI",
		},
		{
			Name: "BAR LA STAZIONE DI ZALLA ALKET",
		},
		{
			Name: "BAR LE PIAGGE DI MECHELLI",
		},
		{
			Name: "BAR LE POSTE S.N.C.",
		},
		{
			Name: "BAR LUNGHINI ANNA",
		},
		{
			Name: "BAR MARO DI ARRIGUCCI ANDREA &",
		},
		{
			Name: "BAR MAYFAIR DI MILLARINI GATTI",
		},
		{
			Name: "BAR MILLENNIO DI MENCARONI",
		},
		{
			Name: "BAR MODERNO S.A.S.DI CAVALLINI",
		},
		{
			Name: "BAR MORGANA DI LO PEZ CAMACHO",
		},
		{
			Name: "BAR NIGHT & DAY DI PATILEA",
		},
		{
			Name: "BAR ORCHIDEA",
		},
		{
			Name: "BAR PASTICCERIA BANCHELLI",
		},
		{
			Name: "BAR PASTICCERIA LA PERLA DI",
		},
		{
			Name: "BAR PASTICCERIA MARCONI DI",
		},
		{
			Name: "BAR PISCINA COMUNALE DI",
		},
		{
			Name: "BAR R.T.L. GIOTTO DI BICHI",
		},
		{
			Name: "BAR RISTORANTE LA NAVE",
		},
		{
			Name: "BAR SHEILA S.N.C.",
		},
		{
			Name: "BAR SORRISO DI GAKHUN ELENA",
		},
		{
			Name: "BAR SPORT CORTONA S.N.C. DI",
		},
		{
			Name: "BAR SUGAR S.A.S DI BURACCHI &C",
		},
		{
			Name: "BAR TRATTORIA TACCONI ANGIOLO",
		},
		{
			Name: "BAR TRE TRE S.R.L.",
		},
		{
			Name: "BAR TROMBETTA BACCHESCHI M.&.C",
		},
		{
			Name: "BAR VECCHIO MULINO",
		},
		{
			Name: "BAR ZAMPAGNI FEDERICO",
		},
		{
			Name: "BAR ZUCCHINI DI ZUCCHINI GUIDO",
		},
		{
			Name: "BARAONDA ROAD BAR DI CAISUTTI",
		},
		{
			Name: "BARBARELLA S.A.S DI FALTONI M&",
		},
		{
			Name: "BARDELLI LEANDRO",
		},
		{
			Name: "BARDI S.N.C DI BARDI ROBERTO",
		},
		{
			Name: "BARNECHI MAURO",
		},
		{
			Name: "BARONI MAURIZIO",
		},
		{
			Name: "BASE ALIMENTARE S.R.L.",
		},
		{
			Name: "BASTIA BEVANDE DISTRIBUZIONE",
		},
		{
			Name: "BASTIAN CONTRARIO DI PIERONI",
		},
		{
			Name: "BASTRENGHI MAURO",
		},
		{
			Name: "BCT SERVICE S.R.L.S",
		},
		{
			Name: "BELIER BIJOURS S. R. L.",
		},
		{
			Name: "BELL'S PUB S.N.C. DI DONNINI E",
		},
		{
			Name: "BEN SLITTANO PUNTO PANE",
		},
		{
			Name: "BENIGNI SERENELLA",
		},
		{
			Name: "BENNATI LORENZA AG.IL GELSO",
		},
		{
			Name: "BERE E SNC DI TIZZANINI S.E CAVIGLI S.",
		},
		{
			Name: "BERNI & RAFFAELLI SNC",
		},
		{
			Name: "BETON PAN S.N.C. DI PANICHI",
		},
		{
			Name: "BETTY OGBONMWAN FAVOUR FROM GO",
		},
		{
			Name: "BEVERYA DI SCOTTO DI LUZIO MASSIMO",
		},
		{
			Name: "BEVILACQUA GABRIELE",
		},
		{
			Name: "BIA 2 S.R.L.",
		},
		{
			Name: "BIAGIANTI DECIO",
		},
		{
			Name: "BIANCHI F. BIANCHI L.",
		},
		{
			Name: "BIBRO' S.R.L.",
		},
		{
			Name: "BICE VIAGGI DI GAZZINI",
		},
		{
			Name: "BIDINI STEFANO EDICOLA IL GIOR",
		},
		{
			Name: "BIRRA PERONI S.R.L",
		},
		{
			Name: "BISTROT CHECCONI DI CHECCONI",
		},
		{
			Name: "BIZZI LUIGI",
		},
		{
			Name: "BLACK PANTHER S.A.S. DI FLORIO",
		},
		{
			Name: "BLETA LEONARD",
		},
		{
			Name: "BLU HOTEL S.R.L VILLA PARADISO",
		},
		{
			Name: "BLU HOTEL S.R.L.HOTEL ABBAZIA",
		},
		{
			Name: "BLU NOTTE S.N.C.",
		},
		{
			Name: "BLUE WATER DI MAGI BRUNO",
		},
		{
			Name: "BLUE WATER S.N.C. DI MAGI",
		},
		{
			Name: "BLUES BAR",
		},
		{
			Name: "BOBI AUTO DI ROBERTO BRILLI",
		},
		{
			Name: "BOCCONDIVINO S.R.L.",
		},
		{
			Name: "BODY LINE S.N.C. DI",
		},
		{
			Name: "BOGHIAN  LACRAMMAIORA",
		},
		{
			Name: "BOLLE & RISPARMIO DI FRAU",
		},
		{
			Name: "BONINSEGNI AUTO S.P.A.",
		},
		{
			Name: "BONTA TOSCANE FORNO LEGNA SNC",
		},
		{
			Name: "BORGO PARADISE CLUB",
		},
		{
			Name: "BORGO RICCIO S:R:L",
		},
		{
			Name: "BOSONE GARDEN S.R.L.",
		},
		{
			Name: "BOTTEGA DELLA BONTA DI BOFFA",
		},
		{
			Name: "BOTTEGA DELLA PASTA FRESCA",
		},
		{
			Name: "BRIGANTI S.R.L.",
		},
		{
			Name: "BRINI MARCO LA GRANDE MELA",
		},
		{
			Name: "BRINI SOC AGR S.S",
		},
		{
			Name: "BRUMEL S.R.L.",
		},
		{
			Name: "BRUNI PIERA",
		},
		{
			Name: "BRUSCHI S.E PAPERINI G. SNC",
		},
		{
			Name: "BUCALETTI FRANCESCO",
		},
		{
			Name: "BUCCI VALERIO",
		},
		{
			Name: "BUON CAFFE S.N.C. DI",
		},
		{
			Name: "BURGER & GRILL DI SULTANALIERA",
		},
		{
			Name: "BUSANEL S.R.L.",
		},
		{
			Name: "BUTOS HO.RE.CA. S.R.L.",
		},
		{
			Name: "BY EAT & MEET S.R.L.S.",
		},
		{
			Name: "C & G DI ELIA CHIATTI E C.",
		},
		{
			Name: "C E S C O T",
		},
		{
			Name: "C.A.E P. GHETTI S.P.A.",
		},
		{
			Name: "C.F.C. S.R.L.",
		},
		{
			Name: "C.G.R DI CASTELLANI GLORIA",
		},
		{
			Name: "C.I.P.A. COMPAGNIA ITALIANA",
		},
		{
			Name: "C.I.S.E.T. DI SALVADORI",
		},
		{
			Name: "C.J.E P.M SNC DI CORNESCHI J E",
		},
		{
			Name: "C.M.I.",
		},
		{
			Name: "C.R FOUD DI VERDELLI & C",
		},
		{
			Name: "CAFE'NOIR S.N.C. DI BACCI A.&",
		},
		{
			Name: "CAFEL ELETTRONICA S.R.L.",
		},
		{
			Name: "CAFFE DEGLI ARTISTI DI",
		},
		{
			Name: "CAFFE DEGLI SPORTIVI DI IACONI",
		},
		{
			Name: "CAFFE DEI ROSSI DAL 1956 S.A.S",
		},
		{
			Name: "CAFFE DEL CORSO",
		},
		{
			Name: "CAFFE DEL TEATRO DI BRUNORI",
		},
		{
			Name: "CAFFE DEL TORREONE DI",
		},
		{
			Name: "CAFFE DEL TRASIMENO DI",
		},
		{
			Name: "CAFFE DELLA SOSTA DI",
		},
		{
			Name: "CAFFE LA COSTA DI SAN ROCCO DI",
		},
		{
			Name: "CAFFE LA MARCHIONNA DI",
		},
		{
			Name: "CAFFE LE MURA DI CONTI SUSANNA",
		},
		{
			Name: "CAFFE LE MURA DI PELUCCHINI",
		},
		{
			Name: "CAFFE MARCO PERENNIO S.A.S.",
		},
		{
			Name: "CAFFE PASTICCERIA DI LUCIANO",
		},
		{
			Name: "CAFFE SAIN 2 DI PUCCI DANIELE",
		},
		{
			Name: "CAFFE SAN DONATO S.N.C. DI",
		},
		{
			Name: "CAFFE SIGNORELLI GE.BA S.N.C.",
		},
		{
			Name: "CAFFE VIA VITTORIO VENETO SRL",
		},
		{
			Name: "CAFFE' DELLA FIERA S.N.C.",
		},
		{
			Name: "CAFFE' VASARI DI MARCHI",
		},
		{
			Name: "CAFFE'69 S.N.C. DI PRESTINICE",
		},
		{
			Name: "CAFFETTERIA VEZZOSI DI SENAPE",
		},
		{
			Name: "CALCAGNO ROSARIO & C.",
		},
		{
			Name: "CALCIT CORTONA VALDICHIANA",
		},
		{
			Name: "CALONI FRANCO",
		},
		{
			Name: "CALOSCI E TOTOBROCCHI S.N.C.",
		},
		{
			Name: "CALVIO PIETRO FORNO BITURGENSE",
		},
		{
			Name: "CALZINI FRANCESCO",
		},
		{
			Name: "CAMBIAMENTI AZIENDALI SRL",
		},
		{
			Name: "CAMET S.N.C.DI PICCIAFUOCHI",
		},
		{
			Name: "CAMST SOC.COOP.A.R.L.",
		},
		{
			Name: "CANCELLONI FOOD SERVICE S.P.A.",
		},
		{
			Name: "CANNELLA CALOGERO & C. SNC",
		},
		{
			Name: "CANOVA DI SAN FRANCESCO DI",
		},
		{
			Name: "CANTARIDI CINZIA",
		},
		{
			Name: "CANTINE BRUSA S.P.A.",
		},
		{
			Name: "CANTINE RIUNITE & CIV",
		},
		{
			Name: "CARNEVALE FOIANO DELLA CHIANA",
		},
		{
			Name: "CASA BELLAVISTA S.N.C. DI",
		},
		{
			Name: "CASA DEL POPOLO SAN LORENZO",
		},
		{
			Name: "CASA DELLA SEDIA DI RONTI",
		},
		{
			Name: "CASA DI RIPOSO SERRISTORI",
		},
		{
			Name: "CASA VINICOLA ZONIN S.P.A.",
		},
		{
			Name: "CASALE CARDINI S.R.L.",
		},
		{
			Name: "CASALE DI BROLIO S.A.S. DI",
		},
		{
			Name: "CASALE L'ANTICO CARRO S.A.S",
		},
		{
			Name: "CASATA DAVINI S.R.L.",
		},
		{
			Name: "CASEIFICIO MATTEASSI ONELIO",
		},
		{
			Name: "CASTELLO DI MODANELLA S.R.L.",
		},
		{
			Name: "CATALDO ALESSIA YOGORINO",
		},
		{
			Name: "CAUTHA ENTE TERZO SETTORE",
		},
		{
			Name: "CDR S.R.L.",
		},
		{
			Name: "CECCARELLI SILVIA ALIMENTARI",
		},
		{
			Name: "CECCONATA GIOVANNI",
		},
		{
			Name: "CENERI SILVANA",
		},
		{
			Name: "CENTO MENU' DI PERUZZI LARA",
		},
		{
			Name: "CENTRO ANZIANI CASTIGLION",
		},
		{
			Name: "CENTRO COMMERCIALE MARINO SRL",
		},
		{
			Name: "CENTRO DELLE CULTURE DI AREZZO",
		},
		{
			Name: "CENTRO DI AGGREGAZIONE SOCIALE",
		},
		{
			Name: "CENTRO FRUTTA DI FICCO MARIA",
		},
		{
			Name: "CENTRO ITALIA SERVIZI SRLS",
		},
		{
			Name: "CENTRO SPORTIVO CORTONA",
		},
		{
			Name: "CENTRO SPORTIVO ITALIANO",
		},
		{
			Name: "CENTROSPESA PETRUCCI S.N.C.",
		},
		{
			Name: "CENTROTTICA S.R.L.",
		},
		{
			Name: "CERES S.P.A.",
		},
		{
			Name: "CHE PIZZA DI FERMENTINI",
		},
		{
			Name: "CHEF AT HOME CORTONA SRLS",
		},
		{
			Name: "CHEF EXSPRES S.P.A.",
		},
		{
			Name: "CHIANUCCI SILVANO",
		},
		{
			Name: "CHIANUCCI VALENTINA",
		},
		{
			Name: "CHIARABOLLI MASSIMO",
		},
		{
			Name: "CHIASSA SPORTING CLUB A.S.D.",
		},
		{
			Name: "CHICKEN TASTE SRLS",
		},
		{
			Name: "CHICKEN TASTE SRLS BOLOGNA",
		},
		{
			Name: "CIAMBELLI ROMANO",
		},
		{
			Name: "CIAO CIAO DI BERNASCONI & C.",
		},
		{
			Name: "CIARDI DECORI DI MENCHETTI",
		},
		{
			Name: "CIPOLLESCHI MASSIMO",
		},
		{
			Name: "CIRCOLO A.S.D. ASLAM",
		},
		{
			Name: "CIRCOLO A.S.P.I ORATORIO",
		},
		{
			Name: "CIRCOLO ACLI DEL RIVAIO",
		},
		{
			Name: "CIRCOLO ACLI LA NOCETA",
		},
		{
			Name: "CIRCOLO ACLI SADAM",
		},
		{
			Name: "CIRCOLO AL DOPO A.P.S.",
		},
		{
			Name: "CIRCOLO ANSPI PIEVE DI CHIO",
		},
		{
			Name: "CIRCOLO ARCI BROLIO",
		},
		{
			Name: "CIRCOLO ARCI CAMUCIA",
		},
		{
			Name: "CIRCOLO ARCI CHIANACCE",
		},
		{
			Name: "CIRCOLO ARCI CIGNANO",
		},
		{
			Name: "CIRCOLO ARCI FARNETA",
		},
		{
			Name: "CIRCOLO ARCI FRASSINETO",
		},
		{
			Name: "CIRCOLO ARCI FRATTICCIOLA",
		},
		{
			Name: "CIRCOLO ARCI JUVENTINA",
		},
		{
			Name: "CIRCOLO ARCI MONTECCHIO",
		},
		{
			Name: "CIRCOLO ARCI OLIVETO",
		},
		{
			Name: "CIRCOLO ARCI OSSAIA",
		},
		{
			Name: "CIRCOLO ARCI RIONE SAN DONATO",
		},
		{
			Name: "CIRCOLO ARCI SAN LORENZO APS",
		},
		{
			Name: "CIRCOLO ARCI VENERE",
		},
		{
			Name: "CIRCOLO C.A.P.IT FERRETTO",
		},
		{
			Name: "CIRCOLO C.S.I. PIETRAIA",
		},
		{
			Name: "CIRCOLO COMBATTENTI ARCI",
		},
		{
			Name: "CIRCOLO CULTURALE BURCINELLA",
		},
		{
			Name: "CIRCOLO CULTURALE CASTRONCELLO",
		},
		{
			Name: "CIRCOLO DI VIA GIOTTO",
		},
		{
			Name: "CIRCOLO ENALS PETRIGNANO",
		},
		{
			Name: "CIRCOLO IL LUPERCALE",
		},
		{
			Name: "CIRCOLO IL QUADRIFOGLIO",
		},
		{
			Name: "CIRCOLO M.C.L. FOIANO",
		},
		{
			Name: "CIRCOLO M.C.L. LA NAVE",
		},
		{
			Name: "CIRCOLO M.C.L. MONTECCHIO",
		},
		{
			Name: "CIRCOLO M.C.L.LA CAPANNINA",
		},
		{
			Name: "CIRCOLO MCL POZZO DELLA CHIANA",
		},
		{
			Name: "CIRCOLO ORATORIO DEL CASSERO",
		},
		{
			Name: "CIRCOLO POKER LIFE TORRITA",
		},
		{
			Name: "CIRCOLO RELAX",
		},
		{
			Name: "CIRCOLO RICREATIVO MONSIGLIOLO",
		},
		{
			Name: "CIRCOLO RIONALE AL CANAPO",
		},
		{
			Name: "CIRCOLO SOCIALE RICREATIVO LA BOCCIOFILA",
		},
		{
			Name: "CIRCOLO SS.MO CROCIFISSO",
		},
		{
			Name: "CIRCOLO TERRITORIALE DEL",
		},
		{
			Name: "CIRELLI VITTORIO LAVORAZIONE",
		},
		{
			Name: "CIRO DI FRANCO PIZZERIA SANTA CROCE",
		},
		{
			Name: "CITTI & CO S.N.C. DI CITTI",
		},
		{
			Name: "CLACAFE' S.R.L. AREA IL GRILLO",
		},
		{
			Name: "CLACAFE' S.R.L. AREA S.PIETRO",
		},
		{
			Name: "CLANIS SNC DI CAPONI D. E",
		},
		{
			Name: "CLASSICA S.R.L.",
		},
		{
			Name: "CLEAN SERVICE S.R.L.",
		},
		{
			Name: "CLF S.P.A.",
		},
		{
			Name: "CLIMACENTER DI CIUCCI MARIO",
		},
		{
			Name: "CLUB VELICO TRASIMENO",
		},
		{
			Name: "CO.ART.E. S.N.C. COSTRUZIONI",
		},
		{
			Name: "CO.GE.DI. INTERNATIONAL S.P.A.",
		},
		{
			Name: "COCA COLA H.B.C. ITALIA S.R.L",
		},
		{
			Name: "COCKTAIL & CO",
		},
		{
			Name: "COCKTAIL MIXOLOGY S.R.L.",
		},
		{
			Name: "COFFEE ART DI SANNO SARA",
		},
		{
			Name: "COLESSEUM TOURS S.R.L.",
		},
		{
			Name: "COMITATO FESTA DI SETTEMBRE",
		},
		{
			Name: "COMITATO FESTEGGIAMEN LA PACE",
		},
		{
			Name: "COMITATO FESTEGGIAMENTI LA PACE",
		},
		{
			Name: "COMITATO ORGOGLIO AMARANTO",
		},
		{
			Name: "COMPAGNIE DEI VIAGGIATORI SRL",
		},
		{
			Name: "COMUNE CASTIGLION FIORENTINO",
		},
		{
			Name: "COMUNE DI CORTONA",
		},
		{
			Name: "COMUS S.R.L.S.",
		},
		{
			Name: "CONFRATERNITA CAMUCIA DELLA",
		},
		{
			Name: "CONFRATERNITA DI MISERICORDIA",
		},
		{
			Name: "CONLEMIEMANI DI BISCONTRI",
		},
		{
			Name: "CONSERVE ITALIA SOC COOP AGR.",
		},
		{
			Name: "CONSIGLIO DEI TERZIERI DI CORTONA",
		},
		{
			Name: "CONSORZIO OPERATRICI TURISTICI",
		},
		{
			Name: "CONSULENTE ENOLOGICA S.R.L.",
		},
		{
			Name: "CONSULTA COMUNALE DEL",
		},
		{
			Name: "CONTI FANCIULLI SRL",
		},
		{
			Name: "CONTI M.E P. & C. SAS",
		},
		{
			Name: "COOP.PRODUZIONE E LAVORO NUOVI",
		},
		{
			Name: "COOPERATIVA-CULTURA-FARNETELLA",
		},
		{
			Name: "CORBELLI ARTURO",
		},
		{
			Name: "CORTONA ACCOGLIENZA S.N.C.",
		},
		{
			Name: "CORTONA CARNI S.A.S DI ANIELLO",
		},
		{
			Name: "CORTONA PADEL SOCIETA SPORTIVA",
		},
		{
			Name: "CORTONA SVILUPPO",
		},
		{
			Name: "CORTONA VACANZE S.R.L.",
		},
		{
			Name: "CORTONA VOLLEY ASSOCIAZIONE",
		},
		{
			Name: "CORTONESE CARNI S.R.L.",
		},
		{
			Name: "COSCI ARGENTINA",
		},
		{
			Name: "COSCI MARCO",
		},
		{
			Name: "CPP CLEMENZ & POHL PRODUCTION",
		},
		{
			Name: "CRAZY BAR CARBONE VLASIE & C.",
		},
		{
			Name: "CREDITO COOPERATIVO DI CHIUSI",
		},
		{
			Name: "CRESCERE ASSOCIAZIONE CULTURALE E DI VOLONTARIATO ODV",
		},
		{
			Name: "CRISTALLO CAFE' S.R.L.",
		},
		{
			Name: "CUCINA SABRINA",
		},
		{
			Name: "CUCULI LUIGI",
		},
		{
			Name: "CUGINI BRUSCHI S.N.C. DI",
		},
		{
			Name: "D.I.A.B. SERVICE S.R.L.",
		},
		{
			Name: "D.M.T. S.R.L. S.U.",
		},
		{
			Name: "D.O.C. DI SCIPIONI FRANCESCO &",
		},
		{
			Name: "D'ORIANO ANTONELLA HOT BAR",
		},
		{
			Name: "D&C STORE S.R.L.",
		},
		{
			Name: "DA GIUDA DI GIORGETTI SIMONA",
		},
		{
			Name: "DA LELLA DI VERSARI ANTONELLA",
		},
		{
			Name: "DA NANNI DI LIVI FRANCESCA R.&",
		},
		{
			Name: "DAL SANTI DI CRISTIAN E DEBORA",
		},
		{
			Name: "DAL TOCIO DISCOBAR DI RENZETTI",
		},
		{
			Name: "DAVIDE CAMPARI-MILANO N.V",
		},
		{
			Name: "DBI SERVICE S.R.L.",
		},
		{
			Name: "DEBITI & C. S.A.S.DEL PIANTA",
		},
		{
			Name: "DEL MONTE FODS SUD EUROPA SPA",
		},
		{
			Name: "DEL VIGNA S.N.C. CAFFE' CHICCO",
		},
		{
			Name: "DELHI KEBAB KING",
		},
		{
			Name: "DELIZIE DELLA VALDICHIANA SRL",
		},
		{
			Name: "DESSOT DI DRINGOLI M. MEACCI G",
		},
		{
			Name: "DI GUSTO DI HUSSAIN WAJID",
		},
		{
			Name: "DI NAVROOP STINGH RANDHAWA",
		},
		{
			Name: "DIAGEO ITALIA S.P.A.",
		},
		{
			Name: "DIDJ S.R.L.",
		},
		{
			Name: "DIEVOLE S.R.L.",
		},
		{
			Name: "DIGIGLIO RAFFAELE",
		},
		{
			Name: "DIMOV NAYDEN LAZAROW",
		},
		{
			Name: "DIONIGI ANDREA FORNO",
		},
		{
			Name: "DIONISO 1989 S.N.C. DI",
		},
		{
			Name: "DISTILLERIE F.LLI CAFFO S.R.L.",
		},
		{
			Name: "DISTRIBUTORE ENI DI FUSARI",
		},
		{
			Name: "DITTA GNERUCCI FRANCO",
		},
		{
			Name: "DITTA GRASSO MARIO",
		},
		{
			Name: "DITTA MONETI MARINO S.A.S.",
		},
		{
			Name: "DITTA QUITTI ROBERTO",
		},
		{
			Name: "DITTA REDI FERNANDO",
		},
		{
			Name: "DITTA REDI FILIPPO",
		},
		{
			Name: "DOBREA NELA FACEBAR",
		},
		{
			Name: "DOGA'S BAR S.N.C.",
		},
		{
			Name: "DOLCE IDEA PALMIERI ROBERTO",
		},
		{
			Name: "DOLCE VITA S.R.L.",
		},
		{
			Name: "DOLCI DELIZIE S.R.L",
		},
		{
			Name: "DOLCI DESIDERI EREDI TIEZZI",
		},
		{
			Name: "DOLCI FOLLIE S.N.C. DI",
		},
		{
			Name: "DOLCI MAGIE S.N.C. DI ALUNNO",
		},
		{
			Name: "DON CARLOS UNIPERSONALE",
		},
		{
			Name: "DOPO LAVORO ELEPHANT CAFE'",
		},
		{
			Name: "DOPPIOZERO DI PAOLUCCI ANNA",
		},
		{
			Name: "DORECA S.P.A.",
		},
		{
			Name: "DOTT.FRANCO QUINTI STUDIO",
		},
		{
			Name: "DOUBLE M S.R.L.",
		},
		{
			Name: "DRAGONI LUCIA ALIMENTARI",
		},
		{
			Name: "DRIVE'IN BAR S.N.C.DI ULIVELLI",
		},
		{
			Name: "DS FITNESS S.R.L.",
		},
		{
			Name: "E.F.A. S.N.C. DI FARALLI",
		},
		{
			Name: "E.VENTO S.R.L.",
		},
		{
			Name: "EDEN S.R.L.",
		},
		{
			Name: "EDILBLOC DI ROSADINI & C. SNC",
		},
		{
			Name: "EFFE 2 S.N.C.",
		},
		{
			Name: "EFFE 4 S.N.C. DI FELICI RENATO",
		},
		{
			Name: "EFFE 5 COSTRUZIONI S.R.L.",
		},
		{
			Name: "EGLE S.R.L.",
		},
		{
			Name: "ELENA MAFUCCI",
		},
		{
			Name: "ELETTROMECCANICA SBRILLI",
		},
		{
			Name: "ELIA CHIATTI &c. s.a.s",
		},
		{
			Name: "ELMAR S.A.S DI SADIKAJ MARINDO",
		},
		{
			Name: "EMC BERTI S.R.L.",
		},
		{
			Name: "EMME2ELLE S.R.L.",
		},
		{
			Name: "EMMETI S.N.C. DI MILANI MARIO",
		},
		{
			Name: "ENEL DISTRIBUZIONE S.P.A",
		},
		{
			Name: "ENGINE S.R.L.",
		},
		{
			Name: "ENOTECA ENOTRIA S.N.C. DI",
		},
		{
			Name: "ENOTRIA S.AS DI SILVIA RINCHI",
		},
		{
			Name: "ENTE FILARMONICO ITALIANO",
		},
		{
			Name: "EREDI FARMACIA BIANCHI",
		},
		{
			Name: "EREDI MILLEMOLLICHE AREZZO",
		},
		{
			Name: "ERIDANIA SADAM S.P.A",
		},
		{
			Name: "ERREEFFE S.R.L",
		},
		{
			Name: "ERREVI DI MINIATI VIVIETTA",
		},
		{
			Name: "ESTETICA CONTATTO FABIANELLI",
		},
		{
			Name: "ETRURIALLARMI S.R.L.",
		},
		{
			Name: "ETRUSCAN SOUL",
		},
		{
			Name: "EURO HALAL FOODS SRLS",
		},
		{
			Name: "EURO VET S.R.L.",
		},
		{
			Name: "EUROAUTO S.R.L.",
		},
		{
			Name: "EUROMECCANICA 2 S.R.L.",
		},
		{
			Name: "EUROSAMAC S.R.L.",
		},
		{
			Name: "EVOLUTION BODY S.R.L.",
		},
		{
			Name: "EXCALIBAR S.A.S. DI BERTI SARA",
		},
		{
			Name: "EXTRABAR CAFFE S.N.C DI G&L",
		},
		{
			Name: "F.A.M. REAL ESTATE S.R.L.",
		},
		{
			Name: "F.C TUSCAN ACADEMY A.S.D",
		},
		{
			Name: "F.C.F. DI FAVILLI FABRIZIO & C",
		},
		{
			Name: "F.G.M. S.A.S DI ROSATI SERENA",
		},
		{
			Name: "F.LLI FABBRI",
		},
		{
			Name: "F.LLI GANCIA & C. S.P.A.",
		},
		{
			Name: "F.LLI PANIZZOLO S.R.L.",
		},
		{
			Name: "F.R. SETTANTA DI FANELLI ROBERTO",
		},
		{
			Name: "F&F GESTIONI S.R.L.",
		},
		{
			Name: "FA.RO S.R.L.",
		},
		{
			Name: "FALTONI CLAUDIO & C. S.R.L.",
		},
		{
			Name: "FAMILY BAR SNC DI LUCA GRASSI",
		},
		{
			Name: "FASCHION CAFE' DI MUSUMECI",
		},
		{
			Name: "FASE 4 S.R.L.",
		},
		{
			Name: "FAST FOOD DONER KEBAB DI IQBAL",
		},
		{
			Name: "FAST FOOD ON THE ROAD DI",
		},
		{
			Name: "FATTORIA \"LE TASSINAIE\"",
		},
		{
			Name: "FEDERAZIONE PROVINCIALE COLDIR",
		},
		{
			Name: "FEDERICO GRAZZINI",
		},
		{
			Name: "FELICI ALDO E FIGLI",
		},
		{
			Name: "FEMAC",
		},
		{
			Name: "FEMAR S.R.L.",
		},
		{
			Name: "FENICE S,R,L",
		},
		{
			Name: "FERRAMENTA LAMBERTI S.N.C. DI",
		},
		{
			Name: "FERRANTI PAOLO",
		},
		{
			Name: "FERRARELLE S.P.A.",
		},
		{
			Name: "FERRERO COMMERCIALE ITALIA SRL",
		},
		{
			Name: "FGS RICAMBI S.R.L.",
		},
		{
			Name: "FICO S.R.L.",
		},
		{
			Name: "FIDELITY SALUS SERVICE &",
		},
		{
			Name: "FILIPPO SABATINI",
		},
		{
			Name: "FIM S.P.A.",
		},
		{
			Name: "FIN.TUR.S.R.L.",
		},
		{
			Name: "FINOCCHI ARCANGIOLO E. FIGLI",
		},
		{
			Name: "FLOOR SERVICE SR.L.",
		},
		{
			Name: "FOFO' S.A.S.",
		},
		{
			Name: "FONDAZIONE MONNALISA ONLUS",
		},
		{
			Name: "FONTANAFREDDA SRL",
		},
		{
			Name: "FONTEMAGGIO SNC DI ARCURI",
		},
		{
			Name: "FOR-BAR S.N.C. DI INNOCENTI",
		},
		{
			Name: "FORDELING LOGISTICS S.R.L.",
		},
		{
			Name: "FORTI ALBERTO",
		},
		{
			Name: "FOTO LAMENTINI DI GEMMI LINA",
		},
		{
			Name: "FP S.R.L.",
		},
		{
			Name: "FRACASSI FABRIZIO ROSTICCERIA",
		},
		{
			Name: "FRANCI KATIA & C. LE COCCINELL",
		},
		{
			Name: "FRANCI S.R.L.",
		},
		{
			Name: "FRANCO MENCI",
		},
		{
			Name: "FRAPPI MARIO S.R.L.",
		},
		{
			Name: "FRAPPI S.R.L.",
		},
		{
			Name: "FRATELLI AVERNA S.P.A.",
		},
		{
			Name: "FRATELLI BRANCA DISTILLERIE",
		},
		{
			Name: "FRATELLI BRUNO COSTRUZIONI SRL",
		},
		{
			Name: "FRATELLI CUCULI S.A.S DI",
		},
		{
			Name: "FRATELLI FALINI S.R.L.",
		},
		{
			Name: "FRATELLI MARTINO DI MARTINO",
		},
		{
			Name: "FRATELLI MENCACCI S.N.C. DI",
		},
		{
			Name: "FRATELLI PIANCIO S.N.C.",
		},
		{
			Name: "FRUTTA E VERDURA DI FUNGHINI",
		},
		{
			Name: "FRUTTA E VERDURA FEGATELLI",
		},
		{
			Name: "FRUTTA E VERDURA LA PIETRA",
		},
		{
			Name: "FRUTTA E VERDURA PICCIAFUOCHI",
		},
		{
			Name: "FRUTTA E VERDURA PIERONI",
		},
		{
			Name: "FRUTTA VERDURA BARBONI CHIARA",
		},
		{
			Name: "FRUTTA VERDURA CHIMERA S.R.L.",
		},
		{
			Name: "FRUTTISSIMA DI NANDESI",
		},
		{
			Name: "FUSARI CARBURANTI DI FUSARI",
		},
		{
			Name: "FUTURA 2 C S.A.S. DI C.V.",
		},
		{
			Name: "FUTURAGRI S.R .L",
		},
		{
			Name: "FUTURE OFFICE S.A.S. DI",
		},
		{
			Name: "G.D.C.GESTIONI DI GUALANO & C.",
		},
		{
			Name: "G.L.IMMOBILIARE S.R.L.",
		},
		{
			Name: "G.M.E. S.R.L.",
		},
		{
			Name: "G.P.ESSE MARCHET",
		},
		{
			Name: "G.R.F. DI PALERMO FRANCESCO &",
		},
		{
			Name: "G.S CONFEZIONI S.R.L",
		},
		{
			Name: "G.S. MANCIANO",
		},
		{
			Name: "G.S. TERONTOLA ASD",
		},
		{
			Name: "G.S.A. S.R.L.(GESTIONI E",
		},
		{
			Name: "G.S.A.D.I. DI GOVERNATORI & C.",
		},
		{
			Name: "GABBELLI S.N.C. DI GABELLI",
		},
		{
			Name: "GALILEO PALACE HOTEL DI EFFE 4",
		},
		{
			Name: "GAMBINO ROBERTO",
		},
		{
			Name: "GARZI ROBERTA",
		},
		{
			Name: "GATTOBIGIO SRLS DI GATTOBIGIO ELISA",
		},
		{
			Name: "GEDYS S.R.L.",
		},
		{
			Name: "GELATERIA COCO PALM DI",
		},
		{
			Name: "GELATERIA DOLCE VITA DI",
		},
		{
			Name: "GELATERIA LEONARDO DI BATTISTI",
		},
		{
			Name: "GELATERIA PRIMO AMORE DI",
		},
		{
			Name: "GELATO TI AMO DI TANVEER SAJID",
		},
		{
			Name: "GEMINI AUTOSERVIZZI DI GEMINI",
		},
		{
			Name: "GEPPONI S.R.L.",
		},
		{
			Name: "GERIST S.R.L. GESTIONE",
		},
		{
			Name: "GESTIONE ALBERGHIERA S.A.S.",
		},
		{
			Name: "GESTIONI S.A.S. DI FATTORINI",
		},
		{
			Name: "GHIOTTO S.R.L.",
		},
		{
			Name: "GIACOMO GIOVAGNINI",
		},
		{
			Name: "GIACOMO KONZ & C. S.P.A.",
		},
		{
			Name: "GIANNINI VILMA",
		},
		{
			Name: "GIESSE MODA S.R.L.",
		},
		{
			Name: "GIL S.A.S. DI BRIGANTI G.& C.",
		},
		{
			Name: "GIOA.MAR DI MARRINI GIORGIA",
		},
		{
			Name: "GIOA.MAR SEVEN       DI",
		},
		{
			Name: "GIOMMETTI S.R.L",
		},
		{
			Name: "GIORGESCHI & C.SNC GELATERIA",
		},
		{
			Name: "GIORGI RICCARDO",
		},
		{
			Name: "GIORI DISTILLATI S.R.L",
		},
		{
			Name: "GIOVANNINI BIBITE DI",
		},
		{
			Name: "GIOVE S.R.L.",
		},
		{
			Name: "GIU.MI SNC DI OBERMAJER CARMEL",
		},
		{
			Name: "GUIDICI RITA  LA TUA PIADINA",
		},
		{
			Name: "GIULIANDREA S.N.C. DI GALLI",
		},
		{
			Name: "GOLDEN HILLS CHAPTER ITALY",
		},
		{
			Name: "GOLF CLUB VALDICHIANA SOC.",
		},
		{
			Name: "GOOD LUCK ONE",
		},
		{
			Name: "GOOD LUCK S.N.C 2",
		},
		{
			Name: "GOOD LUCK S.N.C SERVIZIO",
		},
		{
			Name: "GORELLI S.R.L.",
		},
		{
			Name: "GORI ROBERTO",
		},
		{
			Name: "GOSTINICCHI ALBERTO",
		},
		{
			Name: "GRAFFIO DI LUCA SERENI",
		},
		{
			Name: "GRAN HOTEL SERRE S.R.L",
		},
		{
			Name: "GREEM POINT DI DONNINI MATTEO",
		},
		{
			Name: "GREEN BAR",
		},
		{
			Name: "GREEN GYM FITNESS CENTER A.S.D",
		},
		{
			Name: "GREMOLI FERNANDO",
		},
		{
			Name: "GRUPPO IL CASSERO",
		},
		{
			Name: "GRUPPO ITALIANO VINI S.P.A.",
		},
		{
			Name: "GRUPPO SPORTIVO C.FRATTICCIOLA",
		},
		{
			Name: "GRUPPO SPORTIVO CACCIA AL",
		},
		{
			Name: "GRUPPO SPORTVO CACCIA AL CINGHIALE",
		},
		{
			Name: "GRUPPO STORICO SBANDIERATORI",
		},
		{
			Name: "GRUPPO VINI SELEZIONATI S.R.L",
		},
		{
			Name: "GTA AREZZO S.R.L.",
		},
		{
			Name: "GUIDICCI RITA LA TUA PIADINA",
		},
		{
			Name: "GUSTI SICILIANI S.R.L.S.",
		},
		{
			Name: "GYMNASIUM S.R.L.",
		},
		{
			Name: "HAPPY DAYS DI MIRRA ANTONIETTA",
		},
		{
			Name: "HEINEKEN ITALIA SPA",
		},
		{
			Name: "HI.TECH SPACE",
		},
		{
			Name: "HILL TOWN TOURS S.R.L.",
		},
		{
			Name: "HOSSAIN FORHAD MINI MARKET",
		},
		{
			Name: "HOSTARIA-PIZZERIA \"LA TUFA \"",
		},
		{
			Name: "HOTEL FORUM G.R.A S.R.L.",
		},
		{
			Name: "HOTEL LA VELA S.A.S. DI PEPINI",
		},
		{
			Name: "HOTEL MIRALAGO DI PIESSE SAS",
		},
		{
			Name: "HOTEL SAN LUCA S.N.C.",
		},
		{
			Name: "HOTEL UNAWAY MONTEPULCIANO EST",
		},
		{
			Name: "HUCAIL CRISTINA FLORENTIA",
		},
		{
			Name: "I GUSTOSAPORI ASSOCIAZIONE",
		},
		{
			Name: "I.A.G. DI PAGLIERINI ALFREDO &",
		},
		{
			Name: "ICE BAR S.N.C. DI BINCHI ELISA",
		},
		{
			Name: "IDA IMMOBILIARE S.N.C. DI",
		},
		{
			Name: "IDROGROSS DI COMANDUCCI GIULIO",
		},
		{
			Name: "IGEA S.N.C. DI SCIPIONI A.& S.",
		},
		{
			Name: "IL BORGO S.N.C. DI PROTTI",
		},
		{
			Name: "IL BRIGANTE GENTILUOMO DI MARTINO ANTONIO &C.",
		},
		{
			Name: "IL BUONGUSTAIO DI SBRAGI &",
		},
		{
			Name: "IL BUTTIGHINO S.N.C. DI SONIA",
		},
		{
			Name: "IL CACIO BRILLO DI BUCCI",
		},
		{
			Name: "IL CAPRICCIO DI DELLA GATTA",
		},
		{
			Name: "IL CENTRALE B&B DI FERRI GORI",
		},
		{
			Name: "IL CHIOSCO DELLA STAZIONE",
		},
		{
			Name: "IL FALCONIERE S.R.L.",
		},
		{
			Name: "IL FISCHIO DEL MERLO L.B. SNC",
		},
		{
			Name: "IL FORNAIO PASTICCERE DI",
		},
		{
			Name: "IL GALEONE S.N.C.",
		},
		{
			Name: "IL GAMBERO S.N.C. DI GUERRIENI",
		},
		{
			Name: "IL GATTO & LA VOLPE S.R.L.",
		},
		{
			Name: "IL GERMOGLIO DI LORENZO BASSI",
		},
		{
			Name: "IL GIRASOLE S.N.C. BAR IL",
		},
		{
			Name: "IL GRANDE SOGNO S.R.L.",
		},
		{
			Name: "IL GRIFONE CORTONA RESIDENCE",
		},
		{
			Name: "IL JOLLY S.N.C.",
		},
		{
			Name: "IL MAGGIOLONE DI ALPINI",
		},
		{
			Name: "IL NOCETO DI ASSISI DI MENCI",
		},
		{
			Name: "IL PIACERE DELLA CARNE DI",
		},
		{
			Name: "IL PICCHIO S.R.L",
		},
		{
			Name: "IL PICCOLO FORNAIO SNC DI",
		},
		{
			Name: "IL PONTE MAGICO DI DE ROSA",
		},
		{
			Name: "IL RISTORO DI VIA DANTE",
		},
		{
			Name: "IL SOLCO DI RONDELLI MATTEO",
		},
		{
			Name: "IL SOLE DEL SODO VILLA CLAUDIA",
		},
		{
			Name: "IL TEATRO S:R:L:",
		},
		{
			Name: "IL VECCHIO MULINO DI MAZZONI",
		},
		{
			Name: "IMPRESA EDILE RAGAZZO ENZO",
		},
		{
			Name: "IMPRESA IDILE MORETTI ANDREA",
		},
		{
			Name: "IMPRESA PULIZIE MARIOTTONI SNC",
		},
		{
			Name: "INDUSTRIA DEL GHIACCIO S.N.C.",
		},
		{
			Name: "INSAF MARKET SAS DI AKTER SHRA",
		},
		{
			Name: "INSOMNIA CAFFE' S.N.C. DI",
		},
		{
			Name: "IQBAL NASIR ALIBABA'",
		},
		{
			Name: "ISOLA DEL GELATO DI TELLINI A.",
		},
		{
			Name: "ISTITUTO POVERE FIGLIE DELLE",
		},
		{
			Name: "ISTITUTO STATALE D\"ISTRUZIONE",
		},
		{
			Name: "ISTITUZIONE CULTARE EDUCATIVA",
		},
		{
			Name: "ITALART SERVICE S.R.L.",
		},
		{
			Name: "ITALGIAD SRL",
		},
		{
			Name: "ITALIAN PUB S.N.C.",
		},
		{
			Name: "ITALIAN STYLE SRL",
		},
		{
			Name: "ITALY-MARKET DI APPUHAMY",
		},
		{
			Name: "ITALY-MARKET S.R.L.",
		},
		{
			Name: "ITAM S.P.A.",
		},
		{
			Name: "IZZO RAIMONDO PANETTERIA OLMO",
		},
		{
			Name: "J.L S.A.S DI PICCIO JESSICA",
		},
		{
			Name: "JACOPO DI TONDO NON SOLO",
		},
		{
			Name: "JAGER &PARTNES SRL",
		},
		{
			Name: "JAGER PETRO EGIDIO",
		},
		{
			Name: "JAKER E PARTNES S.R.L.",
		},
		{
			Name: "JAMAL ASLM KHEL",
		},
		{
			Name: "JAZZY CLUB SNC GORMET CAFFE'",
		},
		{
			Name: "JOKER .BAR S.R.L.",
		},
		{
			Name: "JOKER BAR DI BENNATI LUCA",
		},
		{
			Name: "JOLLI ORO S.A.S. DI ROGGI",
		},
		{
			Name: "JSF S.N.C. MANDISON BAR",
		},
		{
			Name: "JUNGHEINRICH ITALIANA",
		},
		{
			Name: "KAMRUL HOSSAIN BD ALIMENTARI",
		},
		{
			Name: "KARIBE CAFE' DI SAMUELI",
		},
		{
			Name: "KAUR SATINDER PAUL REGINA",
		},
		{
			Name: "KHAN ASIF ADEEL",
		},
		{
			Name: "KHAN EAR",
		},
		{
			Name: "KOINE COPERATIVA SOCIALE",
		},
		{
			Name: "L.C.DI LORENZONI EDI & C. SNC",
		},
		{
			Name: "L.P.COMPIUTER S.R.L.",
		},
		{
			Name: "L' ALLEGRA FATTORIA DI LUCANI",
		},
		{
			Name: "L'ANGOLO CAFFE' DI MARIA LAURA",
		},
		{
			Name: "L'ANGOLO DEI SAPORI DI",
		},
		{
			Name: "L'ANGOLO DEI SOGNI DI GUILLARO",
		},
		{
			Name: "L'ANGOLO DEL GOLOSO S.N.C. DI",
		},
		{
			Name: "L'ANTICA DIMORA DI CAMILLONI",
		},
		{
			Name: "L'ANTICA FATTORIA S.R.L",
		},
		{
			Name: "L'ANTICA PIEVE S.R.L.",
		},
		{
			Name: "L'ANTICO FORNO DI MAGI ALBERTO",
		},
		{
			Name: "L'ASSICURESE S.R.L.",
		},
		{
			Name: "L'AURORA SOCIETA COOP.SO.ONLUS",
		},
		{
			Name: "L'ETRUSCO BAR PASTICCERIA DI",
		},
		{
			Name: "L'ORI S.R.L.",
		},
		{
			Name: "L'OSTERIA DI GIOMARELLI S. E",
		},
		{
			Name: "L'ULCERA DEL SIGNOR WILSON",
		},
		{
			Name: "LA BANDITA COUNTRY HOTEL SRLS",
		},
		{
			Name: "LA BOTTEGA DEGLI SFIZI DI LUCA",
		},
		{
			Name: "LA BOTTEGA DEL FUTURO DEI",
		},
		{
			Name: "LA BOTTEGA DEL GUSTO",
		},
		{
			Name: "LA BOTTEGA DEL VINO S.N.C.",
		},
		{
			Name: "LA BOTTEGA DELL'OSTE S.N.C.",
		},
		{
			Name: "LA BOTTEGA DELLA BRACERIA DI",
		},
		{
			Name: "LA BOTTEGA DELLE BONTA",
		},
		{
			Name: "LA BOTTEGA DELLE CHIANACCE DI",
		},
		{
			Name: "LA BOTTEGA DELLE FONTI",
		},
		{
			Name: "LA BOTTEGA DI MERCATALE S.R.L.",
		},
		{
			Name: "LA BOTTEGA DI VIA DARDANO",
		},
		{
			Name: "LA BOTTEGHINA 1962 S.R.L.",
		},
		{
			Name: "LA BOTTEGHINA DI OLMO",
		},
		{
			Name: "LA BOTTEGHINA DI SANTA LUCIA",
		},
		{
			Name: "LA BRACERIA DI MEUCCI EMANUELE",
		},
		{
			Name: "LA CANTINA DELLE PIAGGE DI",
		},
		{
			Name: "LA COMBRICCOLA DI TRAPASSI",
		},
		{
			Name: "LA CORTE DEI PAPI S.A.S.",
		},
		{
			Name: "LA DILIGENZA DI DANIELLO",
		},
		{
			Name: "LA DIMORA CAFFE DI SPENSIERATI",
		},
		{
			Name: "LA DIVINA BOTTEGA DI CHILLERI",
		},
		{
			Name: "LA DOLCE VITA S.N.C.",
		},
		{
			Name: "LA FABBRICA DEL SOLE",
		},
		{
			Name: "LA FENICE SOCIETA S.R.L.",
		},
		{
			Name: "LA FRONTIERA  ASSOCIAZIONE",
		},
		{
			Name: "LA FRUTTA DI MANHATTAN\" DI",
		},
		{
			Name: "LA FRUTTERIA DI ROSSANA DEL",
		},
		{
			Name: "LA FRUTTERIA S.R.L.",
		},
		{
			Name: "LA GRANDE MELA AREZZO",
		},
		{
			Name: "LA LEOPOLDAI PASQUALE CORVINO",
		},
		{
			Name: "LA LOCANDIERA DI PALLINI MARIA",
		},
		{
			Name: "LA PANETTERIA DI GIAIMO",
		},
		{
			Name: "LA PERGOLA S.R.L.",
		},
		{
			Name: "LA PIADINA DI ROSY DI ROSSETTA",
		},
		{
			Name: "LA ROCCA GOURMET S.R.L.S.",
		},
		{
			Name: "LA ROSA BLU DI BUCALETTI MARCO",
		},
		{
			Name: "LA SCUOLA DI CUCINA DI",
		},
		{
			Name: "LA SFIZZERIA S.N.C.",
		},
		{
			Name: "LA SOSTA DEL PRIORE S.A.S",
		},
		{
			Name: "LA SOSTA DI CARLO PAZZAGLIA",
		},
		{
			Name: "LA SPIAGGETTA DI CANONICO",
		},
		{
			Name: "LA TANA DEL BEVITORE DI",
		},
		{
			Name: "LA TANA DEL GUFO",
		},
		{
			Name: "LA TERRAZZA 2 DI BANCHETTI",
		},
		{
			Name: "LA TERRAZZA DI BANCHETTI ANDRE",
		},
		{
			Name: "LA TUA SPESA S.N.C.",
		},
		{
			Name: "LADY'S BAR FLOWERS' BAR S.N.C.",
		},
		{
			Name: "LAMAR S.R.L.",
		},
		{
			Name: "LAMBERTI MARCO RECUPERO MAT.FE",
		},
		{
			Name: "LAMPREGHIOTTO S.R.L.S",
		},
		{
			Name: "LANDI ANNA UN PO' DI FRUTTA",
		},
		{
			Name: "LANDUCCI ALESSANDRO",
		},
		{
			Name: "LAP ANGEL AGENCY SRL",
		},
		{
			Name: "LATTONERIE SAVARELLI GIOVANNI",
		},
		{
			Name: "LAZZERINI RAPPRESENTANZE SRL",
		},
		{
			Name: "LE ANTICHE FORNACI TOSCANE SRL",
		},
		{
			Name: "LE CANAGLIE S.R.L.",
		},
		{
			Name: "LE DELIZIE DEL FORNAIO",
		},
		{
			Name: "LE DELIZIE DI BRAGHETTI LUIGI",
		},
		{
			Name: "LE DOME CAFE' S.R.L.",
		},
		{
			Name: "LE GOLOSITA S.N.C. BAR LE",
		},
		{
			Name: "LE PETIT BRISTRO S:A:S:",
		},
		{
			Name: "LE PRIMIZIE DEL LAGO\"S.A.S.DI",
		},
		{
			Name: "LE RUGHE CORTONA GIULIANTtini",
		},
		{
			Name: "LE SANTUCCIE S.A.S",
		},
		{
			Name: "LETO BARBARA",
		},
		{
			Name: "LIBERTY DI RAMPI MASSIMILIANO",
		},
		{
			Name: "LICA POPESCU S.A.S.",
		},
		{
			Name: "LILLI FRUTTA E VERDURA DI NICA GENICA",
		},
		{
			Name: "LIMONI GIUSEPPE E FRATELLI SNC",
		},
		{
			Name: "LO SCOIATTOLO DI BALDI ALFIO",
		},
		{
			Name: "LO STRAVIZIO S.R.L.",
		},
		{
			Name: "LO STUZZICANGOLO DI FISCHETTI",
		},
		{
			Name: "LO\"SPACCONE,, S.R.L.",
		},
		{
			Name: "LOCANDA I GRIFI DI ALESSANDRA",
		},
		{
			Name: "LODOVICHI RICEVIMENTI S.R.L",
		},
		{
			Name: "LOMBARDIA DI ANNA BALZANO & C.",
		},
		{
			Name: "LOMBRICULTURA CORTONESE DI",
		},
		{
			Name: "LONGINI & BOTARELLI DI LONGINI",
		},
		{
			Name: "LORENZETTI ANDREA",
		},
		{
			Name: "LOSCAFAMILYTATTOO DI BRUNO",
		},
		{
			Name: "LOUNGE & SHOP S.R.L",
		},
		{
			Name: "LUCARONI FRANCESCO",
		},
		{
			Name: "LUCCI ANDREA",
		},
		{
			Name: "LUNA KEBAB & PIZZA SRLS",
		},
		{
			Name: "LUX CAFE' DI LABATE FRANCESCA",
		},
		{
			Name: "M & E S.R.L.",
		},
		{
			Name: "M.A.E.S. DI SONNATI PAOLO & C.",
		},
		{
			Name: "M.B. ALIMENTARI MARCHITTO SAS",
		},
		{
			Name: "M.BELHIMEUR TASTY",
		},
		{
			Name: "M.D.M. S.R.L.",
		},
		{
			Name: "M.G. AUTO CUCULI E TAUCCI",
		},
		{
			Name: "M.G. GALVANICA DI PAGANI",
		},
		{
			Name: "M.T.T. 2003 DI AMENDOLA ANTONI",
		},
		{
			Name: "MA GIA",
		},
		{
			Name: "MA-SER S.N.C.DI SERIACOPI",
		},
		{
			Name: "MA.GI. S.R.L.",
		},
		{
			Name: "MACELLERIA BARBARA DI VIOLA",
		},
		{
			Name: "MACELLERIA CIAMBELLI GUIDO",
		},
		{
			Name: "MACELLERIA MARCELLO DI SARCOLI",
		},
		{
			Name: "MADIP DI DEL CORTO CARLO C.",
		},
		{
			Name: "MAESTRINI S.S.C. DI LUISINI N. ORLANDI D. & C.",
		},
		{
			Name: "MAFALDA DI FRESCUCCI VALERIA",
		},
		{
			Name: "MAFFEIS GUIDO",
		},
		{
			Name: "MAFUCCI SANTI BAR AMICI MIEI",
		},
		{
			Name: "MAGI FRANCESCO",
		},
		{
			Name: "MAGICA GINEVRA DI ROSSI",
		},
		{
			Name: "MAGINI MARCO",
		},
		{
			Name: "MAGLIFICIO S.L. S.n.c",
		},
		{
			Name: "MAJIN BU DI AGGRAVI KRISTIAN",
		},
		{
			Name: "MAKA DISTRIBUZIONE SRL",
		},
		{
			Name: "MALENTACCHI GIAN MAURO",
		},
		{
			Name: "MALIK PIZZA DONER KEBAB",
		},
		{
			Name: "MAMMOOD SULTAN",
		},
		{
			Name: "MANI'S S.R.L.",
		},
		{
			Name: "MANIVA S.P.A",
		},
		{
			Name: "MANZANO SVILUPPO S.R.L",
		},
		{
			Name: "MARCHESI ANTINORI S.P.A.",
		},
		{
			Name: "MARCHESINI DONATELLA",
		},
		{
			Name: "MARCHI MARCELLO",
		},
		{
			Name: "MARCO TAMBURRO",
		},
		{
			Name: "MARES S.R.L.",
		},
		{
			Name: "MARGHERITA PRODOTTI TIPICI",
		},
		{
			Name: "MARGHERITI FOOD & WINE SRL",
		},
		{
			Name: "MARGUTTI GABRIELE",
		},
		{
			Name: "MARRAGHINI CLAUDIO S.N.C.",
		},
		{
			Name: "MARTINI E ROSSI S.P.A",
		},
		{
			Name: "MARYLEBONE BAY S.N.C. DI ENZO",
		},
		{
			Name: "MASSETTI GABRIELLA",
		},
		{
			Name: "MASUM MD JASIM UDDIN",
		},
		{
			Name: "MATE 3 S.N.C. RICAMI",
		},
		{
			Name: "MATERAZZI FABRIZIO",
		},
		{
			Name: "MATERAZZI RICAMI S.R.L.",
		},
		{
			Name: "MATTEO PICCHIONI IL GIRASOLE",
		},
		{
			Name: "MATTI SAMUELE",
		},
		{
			Name: "MAURIZI FRANCO",
		},
		{
			Name: "MAX THEATRE S.R.L.",
		},
		{
			Name: "MAXIM ASSOCIAZIONE CULTURALE",
		},
		{
			Name: "MAYA S.R.L.",
		},
		{
			Name: "MAZZI GABRIELE",
		},
		{
			Name: "MB ELETTRONICA S.R.L.",
		},
		{
			Name: "MD JASHIM UDDIN BANGLA FOOD'S",
		},
		{
			Name: "MECCANICA CECCARELLI & ROSSI",
		},
		{
			Name: "MELEGATTI S.PA.",
		},
		{
			Name: "MENCI & C. S.P.A.",
		},
		{
			Name: "MENCI GIUSEPPE FROLICULTURA",
		},
		{
			Name: "MENCI MARTINA",
		},
		{
			Name: "MENCI PIERGENTINO",
		},
		{
			Name: "MEONI ALVARO S.R.L.",
		},
		{
			Name: "MEONI LORIANA",
		},
		{
			Name: "MERCURIO S.N.C.",
		},
		{
			Name: "MEZZANOTTE ILVA",
		},
		{
			Name: "MG DI GIORGETTI AUTOMAZIONI",
		},
		{
			Name: "MIGLIACCI CLAUDIA FABIOLA",
		},
		{
			Name: "MILANI MACCHINE S.R.L.",
		},
		{
			Name: "MILGHETTI VINICIO",
		},
		{
			Name: "MILIGHETTI CRISTIANO",
		},
		{
			Name: "MILIGHETTI FRATELLI S.N.C. DI",
		},
		{
			Name: "MILK & COFFEE DI CIRELLI E",
		},
		{
			Name: "MILLEMOLLICHE CAMUCIA DI SIMON",
		},
		{
			Name: "MILLENOVECENTO65 DI PALMERINI",
		},
		{
			Name: "MILLOTTI-SACCO S.N.C.",
		},
		{
			Name: "MIMI DI MIMMA DURI",
		},
		{
			Name: "MIMMI DI GUARDABASSI GIOVANNI",
		},
		{
			Name: "MINERALBIRRA S.R.L.",
		},
		{
			Name: "MINI MARKET LUNGHINI",
		},
		{
			Name: "MINI OLIMPIADI 1 CIRCOLO",
		},
		{
			Name: "MINIMARKET LA BOTTEGA DI BADIA",
		},
		{
			Name: "MINUTILLO MAURO",
		},
		{
			Name: "MIO BAR DI MORETTI",
		},
		{
			Name: "MIREA DI BURRONI CHIARA",
		},
		{
			Name: "MIRKO TARTUFI",
		},
		{
			Name: "MIXER S.R.L. ANGOLO CAFFE",
		},
		{
			Name: "MOCA S.N.C. DI CAVALLUCCI &",
		},
		{
			Name: "MOLECOLA SPIRITS SRLS",
		},
		{
			Name: "MOLESINI S.A.S DI MOLESINI",
		},
		{
			Name: "MOLINARI ITALIA S.P.A",
		},
		{
			Name: "MOLINO CARINI DI CARINI",
		},
		{
			Name: "MOLINO PARRI S.R.L.",
		},
		{
			Name: "MOMA' BAR DI MATTEO MONNI",
		},
		{
			Name: "MOMBASA S.N.C. DI VALENTI &",
		},
		{
			Name: "MOMI CAFFE' DI CORBELLI MOIRA",
		},
		{
			Name: "MONDO GELATO DI CASELLI",
		},
		{
			Name: "MONSHI KHAIRUL AMIN",
		},
		{
			Name: "MONTIGIANI PASQUALE",
		},
		{
			Name: "MOON COMUNICAZIONE S.N.C. DI",
		},
		{
			Name: "MORE SANDRO",
		},
		{
			Name: "MORENA DOMENICI",
		},
		{
			Name: "MORETTI FABIO",
		},
		{
			Name: "MORSO & FILETTO S.R.L.",
		},
		{
			Name: "MOTOR SHOP MENGOZZI S.R.L.",
		},
		{
			Name: "MOTORBAR DI ULIVELLI E MORETTI",
		},
		{
			Name: "MOVIDA S.R.L. SEMPLIFICATA",
		},
		{
			Name: "MOVIMAC S.R.L.",
		},
		{
			Name: "MOVIMENTO CRISTIANO LAVORATORI",
		},
		{
			Name: "MOVIMENTO TERRA GRASSO MARIO",
		},
		{
			Name: "MULTISERVICES DI GIULIO",
		},
		{
			Name: "MUNICCHI GIULIANO",
		},
		{
			Name: "MUSEO AI BORGHI ASS. CULTURALE",
		},
		{
			Name: "MY LOVE CAFE DI PIRETTI",
		},
		{
			Name: "NAEEM HOZIFA NEW KEBAB 786",
		},
		{
			Name: "NAMASTEY-INDIA AHMADZAI-ZAR-",
		},
		{
			Name: "NANDESI FRATELLI SDF DI FRANCO",
		},
		{
			Name: "NANNI ANDREA",
		},
		{
			Name: "NAPUL'E' SRL",
		},
		{
			Name: "NARDELLI DONATO & C SNC",
		},
		{
			Name: "NARDINI SERVIZI ALBERGHIERI",
		},
		{
			Name: "NEGRO FRANCESCO",
		},
		{
			Name: "NEMESISS S.R.L.",
		},
		{
			Name: "NEW BAR DI DESTRIERO ANNALISA",
		},
		{
			Name: "NEW BAR ITALY DI GNERUCCI",
		},
		{
			Name: "NEW CENTRO CARTA",
		},
		{
			Name: "NEW GLAMOUR CLUB",
		},
		{
			Name: "NEWQUATTROCAMINI DI ANDREA CERNICCHI",
		},
		{
			Name: "NICE EVENTS S.R.L.",
		},
		{
			Name: "NICOLAE ALEXANDRA",
		},
		{
			Name: "NIGHT CLUB LADY GODIVA DI",
		},
		{
			Name: "NON SOLO FUMO",
		},
		{
			Name: "NONSOLOPANE DI CAPECCHI SERENA",
		},
		{
			Name: "NONSOLOPANE FRATTA DI",
		},
		{
			Name: "NOVELLI TERZILIO",
		},
		{
			Name: "NUCCI UMBELICI VASSELLI S.N.C",
		},
		{
			Name: "NUOVA COOP TORCOLI SOCIETA",
		},
		{
			Name: "NUOVO CIRCOLO RICREATIVO",
		},
		{
			Name: "NUOVO COMITATO FESTEGGIAMENTI",
		},
		{
			Name: "O.A.M. ORDINE ARTI MARZIALI",
		},
		{
			Name: "O.M.A.F.S.N.C. DI FARALLI",
		},
		{
			Name: "OBELIX SOC PRODUZ. SERVIZI",
		},
		{
			Name: "OFFICINA DELLA PIZZA AREZZO",
		},
		{
			Name: "OFFICINA TIEZZI ARDITO",
		},
		{
			Name: "OLIANTI ANGIOLO",
		},
		{
			Name: "OLIMPO S.R.L.",
		},
		{
			Name: "OMAT S.R.L.",
		},
		{
			Name: "ON THE ROAD S.N.C. DI",
		},
		{
			Name: "ON THE ROAD S.R.L.",
		},
		{
			Name: "ONE COFFEE DI MULIU NICOLETTA",
		},
		{
			Name: "ONE S.R.L",
		},
		{
			Name: "OPE.NAZI.ASS.PER C.NAZIONALE",
		},
		{
			Name: "ORATORIO A.N.S.P.I SAN QUIRICO",
		},
		{
			Name: "ORATORIO ANSPI TIBERIADE-",
		},
		{
			Name: "ORATORIO CIRCOLO ANSPI CENTRO",
		},
		{
			Name: "ORATORIO SAN PROSPERO",
		},
		{
			Name: "ORTOFRUTTICOLA LUONGO LUIGI",
		},
		{
			Name: "OSTERIA DEL TEATRO S.R.L.",
		},
		{
			Name: "OSTERIA DELLA MAL'ORA S.R.L.",
		},
		{
			Name: "OSTERIA DELLE TAVERNE DI",
		},
		{
			Name: "OSTERIA DI RE GIOCONDO DI",
		},
		{
			Name: "OTTOCENTRO S.R.L.",
		},
		{
			Name: "OTTORINO PECCHI S.N.C.",
		},
		{
			Name: "OTTOSERVICE S.R.L.",
		},
		{
			Name: "P & P SOCIETA A RESPONSABILITA",
		},
		{
			Name: "P.& G. S.R.L.",
		},
		{
			Name: "P.A.C. DI SPANO",
		},
		{
			Name: "P.S.P. DI PRESENTINI SAURO",
		},
		{
			Name: "P.V. VALERI S.R.L.",
		},
		{
			Name: "PA-RC S.R.L.S.",
		},
		{
			Name: "PACIFICO S.R.L",
		},
		{
			Name: "PAGUS S.R.L.",
		},
		{
			Name: "PALACE S.R.L.",
		},
		{
			Name: "PAMCO S.R.L.",
		},
		{
			Name: "PANETTERIA AGILLA S.A.S. DI",
		},
		{
			Name: "PANETTERIA AURORA DI CONTEMORI",
		},
		{
			Name: "PANETTERIA LA SOSTA CHE GUSTA",
		},
		{
			Name: "PANICHI ALDO",
		},
		{
			Name: "PANIFICIO CASALINGO F.LLI",
		},
		{
			Name: "PANIFICIO CORTONESE s.r.l",
		},
		{
			Name: "PANIFICIO FRATELLI LAZZERI SNC",
		},
		{
			Name: "PANIFICIO FRATELLI MEZZETTI",
		},
		{
			Name: "PANIFICIO LA CIACCIA S.N.C.",
		},
		{
			Name: "PANIFICIO LA TORRE DI AGUTOLI",
		},
		{
			Name: "PANIFICIO LODOVICHI ANGIOLO",
		},
		{
			Name: "PANIFICIO SANTICCIOLI DOMENICO",
		},
		{
			Name: "PANIFICIO SEGANTINI R. &",
		},
		{
			Name: "PANINERIA DAL LEMBO",
		},
		{
			Name: "PANINOTECA DAL BUMBA",
		},
		{
			Name: "PANINOTECA MARI E MONTI DI",
		},
		{
			Name: "PAOLELLI  ALESSANDRO IMPIANTI",
		},
		{
			Name: "PAOLONI & LUNGHINI",
		},
		{
			Name: "PAPPA E CENA S.R.L.",
		},
		{
			Name: "PARMALAT S.P.A.",
		},
		{
			Name: "PARROCCHIA S.G.EVANGELISTA",
		},
		{
			Name: "PARTESA CENTRO S.R.L.",
		},
		{
			Name: "PARTITO DELLA RIFONDAZIONE",
		},
		{
			Name: "PARTITO DEMOCRATICO",
		},
		{
			Name: "PARTITO DEMOCRATICO CIRCOLO BETTOLLE",
		},
		{
			Name: "PARTITO DEMOCRATICO CORTONA",
		},
		{
			Name: "PARTITO DEMOCRATICO SINISTRA",
		},
		{
			Name: "PASTICCERIA  FANTASY",
		},
		{
			Name: "PASTICCERIA ARETINA DI SGREVI",
		},
		{
			Name: "PASTICCERIA BANCHELLI S.A.S",
		},
		{
			Name: "PASTICCERIA BRUSCHI FRATELLI",
		},
		{
			Name: "PASTICCERIA DOLCE VITA SRLS",
		},
		{
			Name: "PASTICCERIA MENCI MAURO",
		},
		{
			Name: "PASTIFICIO FABIANELLI S.P.A",
		},
		{
			Name: "PATWARY SHARAFAT INSAF MARKET",
		},
		{
			Name: "PECCATI DI GOLA DI GROSSO",
		},
		{
			Name: "PELLEGRINI S.P.A.",
		},
		{
			Name: "PENSIERI GRANDI E PICCOLI DI",
		},
		{
			Name: "PEPE NERO S.N.C. DI GIANNINI",
		},
		{
			Name: "PERENNIO CARBURANTE S.N.C.",
		},
		{
			Name: "PERNIGOTTI S.P.A.",
		},
		{
			Name: "PERNOD RICARD ITALIA S.P.A.",
		},
		{
			Name: "PERUGIA CHAPTER",
		},
		{
			Name: "PESCHERIA IL BUONO DI MAZARA",
		},
		{
			Name: "PESCHERIA IL CANTINONE S.N.C.",
		},
		{
			Name: "PESCHERIA MARE-BLU",
		},
		{
			Name: "PESCHERIA MARIMAR S.R.L.",
		},
		{
			Name: "PETRO EGIDIO JAGER",
		},
		{
			Name: "PIADINERIA SAS DI ALTIERI",
		},
		{
			Name: "PIANO B DID GALLUS MARLENA JAN",
		},
		{
			Name: "PIAZZA DEGLI ARTISTI S.R.L.",
		},
		{
			Name: "PIEMME DI MELCANTINI PAOLO",
		},
		{
			Name: "PIERINI BRUNO",
		},
		{
			Name: "PIERINI DINO",
		},
		{
			Name: "PIERLE S.R.L.",
		},
		{
			Name: "PIEROZZI AUTO DI PIEROZZI",
		},
		{
			Name: "PIEROZZI SILVANO",
		},
		{
			Name: "PIGY FAMILI",
		},
		{
			Name: "PIRAMIDE S.N.C BAR APPY DAYS",
		},
		{
			Name: "PISCINA CORTONA TENNIS CLUB",
		},
		{
			Name: "PIZZA & POI..DI MURANO.COPPOLA",
		},
		{
			Name: "PIZZA E FOCACCIA DI CARRETTI",
		},
		{
			Name: "PIZZA MIA DA SALVATORE DI",
		},
		{
			Name: "PIZZERIA ANEMA E CORE",
		},
		{
			Name: "PIZZERIA BAR CHIARO SCURO SRLS",
		},
		{
			Name: "PIZZERIA CROCE DEL TRAVAGLIO",
		},
		{
			Name: "PIZZERIA GINA CAMAR S.R.L.",
		},
		{
			Name: "PIZZERIA LE ROCCHE DI GIOIA",
		},
		{
			Name: "PIZZERIA LUPETTI DI LUPETTI",
		},
		{
			Name: "PIZZERIA MANI D'ORO DI",
		},
		{
			Name: "PIZZERIA PORCINI DI PORCINI",
		},
		{
			Name: "PIZZERIA SAN LAZZERO",
		},
		{
			Name: "PIZZERIA ZIO PEPE DI PONTIERI",
		},
		{
			Name: "PIZZIAMOCI SNC DI HYKA FLORIND",
		},
		{
			Name: "PIZZICOTTO DI BUONO",
		},
		{
			Name: "PIZZIDEA SANCK G.&C. GIULIANI",
		},
		{
			Name: "PLANET BAR S.N.C. DI ROSSI E",
		},
		{
			Name: "PODERE SAN SANINO- MONTEFOL",
		},
		{
			Name: "POGGIO SANT'ANGELO S.N.C. DI",
		},
		{
			Name: "POKER BAR SAFIN ADRIANA&ELENA",
		},
		{
			Name: "POLDO DI DE BIANCHI YARI",
		},
		{
			Name: "POLEZZI GINO",
		},
		{
			Name: "POLIS SOC. COOP SOCIALE",
		},
		{
			Name: "POLISPO.MONTECCHIO VESPONI",
		},
		{
			Name: "POLISPORTIVA  DILETTANTISTICA",
		},
		{
			Name: "POLISPORTIVA CORITO FREESPORT",
		},
		{
			Name: "POLISPORTIVA DILETT.PIEVE AL",
		},
		{
			Name: "POLISPORTIVA FARNETELLA",
		},
		{
			Name: "POLISPORTIVA MONTECCHIO",
		},
		{
			Name: "POLISPORTIVA PARTERRE",
		},
		{
			Name: "POLISPOSTIVA VAL DI LORETO",
		},
		{
			Name: "PRE.FE.LE. S.R.L.",
		},
		{
			Name: "PREMIO INTERNAZIONALE SEMPLICE",
		},
		{
			Name: "PRESENTINI MARTINA",
		},
		{
			Name: "PRESENTINI PAOLO",
		},
		{
			Name: "PRESSLAB ASSOCIAZIONE",
		},
		{
			Name: "PRESTIGE BAR DI BERTI PAOLO",
		},
		{
			Name: "PRESTIGE HOTEL AREZZO S.R.L.",
		},
		{
			Name: "PRIMIZIE",
		},
		{
			Name: "PRO LOCO TEVERINA DI CORTONA",
		},
		{
			Name: "PRO-LOCO DI LISCIANO NICCONE",
		},
		{
			Name: "PRO-LOCO SAGRA PESCE",
		},
		{
			Name: "PROGETTI S.R.L.",
		},
		{
			Name: "PROJECT 18 S.R.L.",
		},
		{
			Name: "PROLOCO CASTIGLION FIORENTINO",
		},
		{
			Name: "PROLOCO FARNETA",
		},
		{
			Name: "PROMOZIONE IMMOBILIARE",
		},
		{
			Name: "PROSA MAURIZIO",
		},
		{
			Name: "PROSPETTIVA SOCIETA COPERATIVA",
		},
		{
			Name: "PUB NEW FUN S.R.L.",
		},
		{
			Name: "PUB PIZZERIA QUO VADIS S.N.C.",
		},
		{
			Name: "PUB THE LIONS WELL",
		},
		{
			Name: "PUBLIC POIUNTI S.N.C.",
		},
		{
			Name: "PUNJAB TANDOORI S.N.C.",
		},
		{
			Name: "PUNTO PIZZA OUR E CHOICE DI",
		},
		{
			Name: "QUADRIFOGLIO S.P.A",
		},
		{
			Name: "QUENCH ACQUAEVINO SRL",
		},
		{
			Name: "QUINTI S.R.L.",
		},
		{
			Name: "R.D. ORO S.N.C. DI SESTINI",
		},
		{
			Name: "RAGAZZI SPECIALI ASSOCIAZIONE",
		},
		{
			Name: "RAJA BAZAR DI RAJA MASHAM",
		},
		{
			Name: "RAJA WAJID HUSSAIN",
		},
		{
			Name: "RASHID RUMANA NEEDS BAZAR",
		},
		{
			Name: "RATRI WAHAD BD SHOP",
		},
		{
			Name: "RAVINDU ASHAN UDUGAMPALA",
		},
		{
			Name: "RE.VI. UFFICIO SNC DI MANGANI",
		},
		{
			Name: "REALISE MUSIC SRLS",
		},
		{
			Name: "RED BULL S.R.L",
		},
		{
			Name: "REDI FILIPPO",
		},
		{
			Name: "REGIRO' S.A.S.",
		},
		{
			Name: "REINHARD THOMAS ANDREAS",
		},
		{
			Name: "RELAIS LA SOLAIA",
		},
		{
			Name: "RELAIS VILLA PETRISCHIO S.R.L.",
		},
		{
			Name: "RENATA EVENTI DI BIANCHI",
		},
		{
			Name: "RESIDENCE IL CASALE S.R.L.",
		},
		{
			Name: "RESIDENCE LE CORNIOLE RETAGGIO",
		},
		{
			Name: "RESIDENZA IL CORSO SOCIETA",
		},
		{
			Name: "RESIDENZE RIABILITATIVE SOC.",
		},
		{
			Name: "RIAZ KHURRAM RIAZ KEBAB",
		},
		{
			Name: "RICCI MINIATI DI RICCI BRUNO &",
		},
		{
			Name: "RICEVIMENTI GRAZIANI & BEGNINI",
		},
		{
			Name: "RIFONDAZIONE COMUNISTA AREZZO",
		},
		{
			Name: "RIMAT SR.L.",
		},
		{
			Name: "RIMINI COOP SOCIETA COOPERATIV",
		},
		{
			Name: "RINCHI FRANCESCA",
		},
		{
			Name: "RIONE COLONNA APS",
		},
		{
			Name: "RIONE PORTA ROMANA",
		},
		{
			Name: "RIST. LA LOCANDA DEL CASTELLO",
		},
		{
			Name: "RIST. MILU S.R.L",
		},
		{
			Name: "RISTOBAR PRIMO BINARIO DI",
		},
		{
			Name: "RISTORANTE \"DA LUCIANO\"",
		},
		{
			Name: "RISTORANTE \"TONINO\" S.R.L.",
		},
		{
			Name: "RISTORANTE A'LIVELLA S.R.L.",
		},
		{
			Name: "RISTORANTE AL TOCCO",
		},
		{
			Name: "RISTORANTE ALBEROTONDO DI",
		},
		{
			Name: "RISTORANTE AMBROSIA DI SCIARRI",
		},
		{
			Name: "RISTORANTE ANTICA DOGANA DI",
		},
		{
			Name: "RISTORANTE BAR PIZZERIA DA LEO",
		},
		{
			Name: "RISTORANTE CINESE FENG SHOU",
		},
		{
			Name: "RISTORANTE DA MATTEO DE WAELE",
		},
		{
			Name: "RISTORANTE ENOTECA OPERA",
		},
		{
			Name: "RISTORANTE GIANO DI DEVIS",
		},
		{
			Name: "RISTORANTE I FARAGLIONI S.N.C",
		},
		{
			Name: "RISTORANTE IL GROTTINO DI",
		},
		{
			Name: "RISTORANTE IL PASSO DI GIANO",
		},
		{
			Name: "RISTORANTE IL PESCATORE DI",
		},
		{
			Name: "RISTORANTE IL PRELUDIO S.R.L",
		},
		{
			Name: "RISTORANTE LA BADIACCIA",
		},
		{
			Name: "RISTORANTE LA CAPANNINA DI",
		},
		{
			Name: "RISTORANTE LA FOCE S.R.L.",
		},
		{
			Name: "RISTORANTE LA LOCANDA DEL",
		},
		{
			Name: "RISTORANTE LA TAVERNA DI JULIO",
		},
		{
			Name: "RISTORANTE MATTEO SARAPE",
		},
		{
			Name: "RISTORANTE PIZZE.IL PASSAGIO",
		},
		{
			Name: "RISTORANTE PIZZERIA I PECCATI",
		},
		{
			Name: "RISTORANTE PIZZERIA IL GALLO",
		},
		{
			Name: "RISTORANTE PIZZERIA PUNTO 16",
		},
		{
			Name: "RISTORANTE PIZZERIA SOTTO AL",
		},
		{
			Name: "RISTORANTE ROGGI S.N.C. DI",
		},
		{
			Name: "RISTORANTE SAN LAZZARO DI",
		},
		{
			Name: "RISTORBAR QUAGLIA S.A.S.",
		},
		{
			Name: "RISTORO PUNTABELLA DI",
		},
		{
			Name: "RIZVANI LAURA",
		},
		{
			Name: "RO.MI DI FANELLI ROBERTO",
		},
		{
			Name: "RODRIGUEZ ACOSTA JAQUELINE",
		},
		{
			Name: "ROGHI ROBERTO",
		},
		{
			Name: "ROMITI FEDERICO",
		},
		{
			Name: "ROSADINI ANDREA",
		},
		{
			Name: "ROSADINI GIUSEPPE",
		},
		{
			Name: "ROSATIQ S.R.L.",
		},
		{
			Name: "ROSSI DANIELA IL FORNAIO",
		},
		{
			Name: "ROSSI MORENO L'ANGOLO GIUSTO",
		},
		{
			Name: "ROSSOW CATERING GMBH",
		},
		{
			Name: "ROXI BAR DI PANARESE ANDREA",
		},
		{
			Name: "RUBECHINI BEATRICE & C.",
		},
		{
			Name: "RUGBY CLANIS CORTONA ASD",
		},
		{
			Name: "RUSCONI GIUSEPPE",
		},
		{
			Name: "RUSH88 DI IURI ANSELMI",
		},
		{
			Name: "RUSSO ERMENEGILDO",
		},
		{
			Name: "S.A.L.T.U. S.R.L.",
		},
		{
			Name: "S.A.S FATTORIA PALAZZO VECCHIO",
		},
		{
			Name: "S.AMANTE STAZIONE ESSO",
		},
		{
			Name: "S.M. SNC DI SCRUDATO MANUELE C",
		},
		{
			Name: "S.P.S. SRL",
		},
		{
			Name: "S.T.A.C. S.R.L. BAR HALLOWEEN",
		},
		{
			Name: "SA.OR. S.R.L. DI CIRELLI",
		},
		{
			Name: "SABI",
		},
		{
			Name: "SACCONE ENZO",
		},
		{
			Name: "SAE S.N.C. DI ROSSI & LUCONI",
		},
		{
			Name: "SAGI S.R.L.",
		},
		{
			Name: "SAGRA DELLA RANOCCHIA CHIANINA",
		},
		{
			Name: "SAKURA JIM JINGZHU",
		},
		{
			Name: "SAL-CO S.N.C. DI COLZI ATTILIO",
		},
		{
			Name: "SALTALBERO SSD ARL",
		},
		{
			Name: "SALVADORI ROMOLO S.A.S DI",
		},
		{
			Name: "SAN BERNARDO S.P.A",
		},
		{
			Name: "SAN GIOVANNI EVANGELISTA",
		},
		{
			Name: "SAN PELLEGRINO S.P.A.",
		},
		{
			Name: "SAN PELLEGRINO S.P.A. *RISTOP*",
		},
		{
			Name: "SAN PELLEGRINO S.P.A.*TEAM P.*",
		},
		{
			Name: "SANCHO PANZA DI MUNETON MANUEL",
		},
		{
			Name: "SANPELLEGRINO SPA RETE HORECA",
		},
		{
			Name: "SAPORE DI PANE E POI DI ROSSI",
		},
		{
			Name: "SARI S.R.L. DAL MORO",
		},
		{
			Name: "SARTU' MICHELINA GALLORANO",
		},
		{
			Name: "SBANCHI IVO",
		},
		{
			Name: "SBIETTI ASSUNT DI LANINI PAOLO",
		},
		{
			Name: "SCIATU MIO DI LORENZO GENNA",
		},
		{
			Name: "SCUOLA DELL'INFANZIA MARIA",
		},
		{
			Name: "SCUOLA MATERNA SAN BIAGIO",
		},
		{
			Name: "SCUOLA VETERINARIA OMEOPATICA",
		},
		{
			Name: "SEBASTIANI EDOARDO",
		},
		{
			Name: "SEBASTIANI GEOMETRA EDOARDO",
		},
		{
			Name: "SEGANTINI ERNESTO",
		},
		{
			Name: "SEIZETA DI IZZO VINCENZO",
		},
		{
			Name: "SELEZIONE AUTO S.R.L",
		},
		{
			Name: "SELICA S.R.L",
		},
		{
			Name: "SERAFINI VALERIO",
		},
		{
			Name: "SERENA WINES 1881 SRL",
		},
		{
			Name: "SERENISSIMA RISTORAZIONE S.P.A",
		},
		{
			Name: "SERRISTORI 2000 S.R.L.",
		},
		{
			Name: "SEVEN DI PARNETTI C.& C. S.A.S",
		},
		{
			Name: "SFERE S.R.L.",
		},
		{
			Name: "SHINE GROUP SRL",
		},
		{
			Name: "SHO TRASPORT E00D",
		},
		{
			Name: "SHOP24 DI TUFA JUSTINA",
		},
		{
			Name: "SHOPPING DEL BORGHETTO DI",
		},
		{
			Name: "SHOPPING DELLA FRUTTA DI",
		},
		{
			Name: "SICUREZZA ASSISTENZA DI",
		},
		{
			Name: "SILLA'S BISTRO DI BROCCHI",
		},
		{
			Name: "SKILTA SPA SB STP",
		},
		{
			Name: "SLURP S.R.L.",
		},
		{
			Name: "SMALL WONDER ITALY S.R.L:",
		},
		{
			Name: "SNACK BAR DI GALLORINI & C.",
		},
		{
			Name: "SNACK BAR PIETRAIA DI",
		},
		{
			Name: "SO.GE.PU. S.P.A.",
		},
		{
			Name: "SOCIETA  AGRICOLA  LAGRIS S.R.L.",
		},
		{
			Name: "SOCIETA AG.RESIDENCE IL BORGO",
		},
		{
			Name: "SOCIETA AGRICOLA BADIA DI",
		},
		{
			Name: "SOCIETA AGRICOLA FRAPPI GINO &",
		},
		{
			Name: "SOCIETA AGRICOLA INNOVATION LA",
		},
		{
			Name: "SOCIETA AGRICOLA LA BUCA S.S.",
		},
		{
			Name: "SOCIETA AGRICOLA LA QUERCE",
		},
		{
			Name: "SOCIETA AGRICOLA LA SCURE SAS",
		},
		{
			Name: "SOCIETA AGRICOLA VALENTINI S.S",
		},
		{
			Name: "SOCIETA AGRICOLAFLORICOLA FLLI",
		},
		{
			Name: "SOCIETA AGROALIMENTARE VALTIBE",
		},
		{
			Name: "SOCIETA COOP AGRICOLA ARL",
		},
		{
			Name: "SOCIETA COOPERATIVA AGRIZOO",
		},
		{
			Name: "SOCIETA FRATELLANZA ARTIGIANA",
		},
		{
			Name: "SOCIETA ITALIANA COSTRUZIONI",
		},
		{
			Name: "SOCIETA' DILETTANTISTICA RUZZO",
		},
		{
			Name: "SODALIZIO PER LA PASTORALE",
		},
		{
			Name: "SOGEGROSS S.P.A.",
		},
		{
			Name: "SOLE LUNA TABACCHERIA DI",
		},
		{
			Name: "SOLELUNI OFFICINA VIDEO DI",
		},
		{
			Name: "SOLFANELLI ILVA",
		},
		{
			Name: "SOLIDARIETA AUSER CAMUCIA",
		},
		{
			Name: "SORGENTE TESORINO S.P.A",
		},
		{
			Name: "SORRENTINO S.R.L",
		},
		{
			Name: "SORRISO S.R.L",
		},
		{
			Name: "SOTTOPIZZA KEBAB E PIZZA",
		},
		{
			Name: "SOTTOVOCE DI MORINI CHIARA",
		},
		{
			Name: "SOTTOVOCE S.R.L.",
		},
		{
			Name: "SOUL BAR CAFFE' DI BOTTI",
		},
		{
			Name: "SPIZZICHI E  BOCCONI DI VALENTINA DI LORENZO &C.",
		},
		{
			Name: "SPORTCLUB CASTIGLIONFIORENTINO",
		},
		{
			Name: "SPORTS EVENTS AICS ASD",
		},
		{
			Name: "STANGANINI EBE S.N.C. DI MEONI",
		},
		{
			Name: "STAR KEBAB DI KHAN MURAD",
		},
		{
			Name: "STAR ONE S.A.S. DI FANCIULLI",
		},
		{
			Name: "STAZIONE SERVIZIO AGIP CAFE'",
		},
		{
			Name: "STAZIONE SERVIZIO ESSO F.B.DI",
		},
		{
			Name: "STEFANO SALANI",
		},
		{
			Name: "STERLIZTRE S.R.L.",
		},
		{
			Name: "STOCK S.R.L.",
		},
		{
			Name: "STRINO ANNAMARIA PIZZERIASANTA CROCE",
		},
		{
			Name: "STUDIO COMMERCIALE BASSINI &",
		},
		{
			Name: "STUDIO GEOMETRI ASSOCIATI",
		},
		{
			Name: "STUDIO GRAPHIK PM S.R.L.",
		},
		{
			Name: "STUDIO TRIBUTARIO BRACCIALI",
		},
		{
			Name: "SUITE SAN MICHELE S.R.L. S.",
		},
		{
			Name: "SUPERCAMPING SAS DI DESIDERI",
		},
		{
			Name: "SWEET REVENGE BAR S.N.C.",
		},
		{
			Name: "T.C.OR.  S.R.L.",
		},
		{
			Name: "TABACCHERIA CALONI GRAZZIELLA",
		},
		{
			Name: "TABACCHERIA CASUCCI",
		},
		{
			Name: "TABACCHERIA FRANCESCHINI MAURO",
		},
		{
			Name: "TABACCHERIA N.13 DONATI",
		},
		{
			Name: "TABACCHERIA ROSMINI S.R.L.",
		},
		{
			Name: "TABACCHERIA SARA DI SARA",
		},
		{
			Name: "TABACCHERIA VITI DI VITI LUCA",
		},
		{
			Name: "TABACCHERIA ZACCHEI GIORGIO",
		},
		{
			Name: "TABAKU MEJREME",
		},
		{
			Name: "TAMAGNINI VERDIANA",
		},
		{
			Name: "TAMBURINI S.R.L.",
		},
		{
			Name: "TANGANELLI ORESTE AUTOTRASPORT",
		},
		{
			Name: "TANGANELLI REMO LAVORAZIONE",
		},
		{
			Name: "TANGANELLI SANTI",
		},
		{
			Name: "TANI GIUSEPPE DI TANI D. & C.",
		},
		{
			Name: "TAPPEZZERIA PALARCHI ANDREINO",
		},
		{
			Name: "TARIQ KEBAB DI AYYOB ARHAM",
		},
		{
			Name: "TATIALE FILMS INC. RAPP. IN",
		},
		{
			Name: "TAVANTI B. S.N.C. DI LODOVICHI",
		},
		{
			Name: "TAVANTI MARINA L'ANTICA PIAZZA",
		},
		{
			Name: "TEAM G&S CORSE",
		},
		{
			Name: "TENIMENTI LUIGI D'ALESSANDRO",
		},
		{
			Name: "TENNIS CLUB CASTIGLIONESE ASD",
		},
		{
			Name: "TENNIS CLUB CORTONA",
		},
		{
			Name: "TENUTA AGRICOLA TORRE A CENAIA",
		},
		{
			Name: "TENUTA BRANCOLETA DI",
		},
		{
			Name: "TENUTE DEL CERRO S.P.A.",
		},
		{
			Name: "TERRETRUSCHE S.R.L.",
		},
		{
			Name: "TERZAROLI MIRCO IMPRESA",
		},
		{
			Name: "TERZIERI PORTA FIORENTINA",
		},
		{
			Name: "THE FUTURE S.R.L.",
		},
		{
			Name: "THE LOCAL BITE DI BANARAS ALI",
		},
		{
			Name: "TI.SA.57 S.R.L.",
		},
		{
			Name: "TICABI S.N.C",
		},
		{
			Name: "TIEMME S.P.A.",
		},
		{
			Name: "TIEZZI PAOLO",
		},
		{
			Name: "TIN S.R.L.",
		},
		{
			Name: "TIN TIN CAFFE' DI KAMINSKA",
		},
		{
			Name: "TIN-LOGISTIC S.R.L",
		},
		{
			Name: "TIPICIDOC S.R.L.",
		},
		{
			Name: "TIRAMISU' PASTICCERIA ARTIGIAN",
		},
		{
			Name: "TOGETHER JEWELS S.R.L.",
		},
		{
			Name: "TOMAIFICIO LINEA DI BRILLI",
		},
		{
			Name: "TOMMASO PONS",
		},
		{
			Name: "TOR CERVARA S.P.A",
		},
		{
			Name: "TORELLI GIANFRANCO",
		},
		{
			Name: "TOSCANA CEREALI SOC.COOP AGR.",
		},
		{
			Name: "TOSCANA IMMAGINE MODA S.R.L.",
		},
		{
			Name: "TOSO S.P.A.",
		},
		{
			Name: "TRADE LOGISTIC NETWORK S.R.L",
		},
		{
			Name: "TRAIL RUNNING PROJECT",
		},
		{
			Name: "TRATTORIA CASTEL GIRARDI",
		},
		{
			Name: "TRATTORIA DEL FORNO",
		},
		{
			Name: "TRAVERTINO S.ANDREA S.R.L.",
		},
		{
			Name: "TRAVERTINO TOSCANO S.P.A.",
		},
		{
			Name: "TRE DI MENNI ALESSANDRO",
		},
		{
			Name: "TREMORI RENATO S.R.L.",
		},
		{
			Name: "TRINACRIA FOOD ANGELO MURANA",
		},
		{
			Name: "TRONTI STEFANO",
		},
		{
			Name: "TUSCANY SUITE PINZOTTO",
		},
		{
			Name: "TUSCANY& PARTENERS LOCANDA.B43",
		},
		{
			Name: "U.C.SINALUNGHESE",
		},
		{
			Name: "U.P.D.TUORO",
		},
		{
			Name: "UCI CENTRO S.R.L.A S.U.",
		},
		{
			Name: "ULMI CRISTIAN CAFFE DELL'UNA",
		},
		{
			Name: "UMAMI S.R.L.",
		},
		{
			Name: "UMASS S.R.L.",
		},
		{
			Name: "UMBELICI B.E VASELLI V. S.N.C.",
		},
		{
			Name: "UNICOM S.A.S",
		},
		{
			Name: "UNICOSI S.N.C. DI ULIVI NICOLA",
		},
		{
			Name: "UNIONE POL. DI.VALDIPIERLE",
		},
		{
			Name: "UNIONE POLISPORTIVA TUORO",
		},
		{
			Name: "UNIONE SPORTIVA CASTIGLIONESE",
		},
		{
			Name: "UNIVERSITI PIZZA DI MAURO",
		},
		{
			Name: "USD VIRTUS LIGNANO",
		},
		{
			Name: "USIL CURTUN S.R.L.",
		},
		{
			Name: "USMAN RIAZ TAJ MAHAL",
		},
		{
			Name: "V.A.B.VIGILANZA ANTINCENDI",
		},
		{
			Name: "V.C.E. S.R.L.",
		},
		{
			Name: "V.O.G.  S.R.L.",
		},
		{
			Name: "VALBONESI CATIA RIV.TABACCHI",
		},
		{
			Name: "VALDICHIANA CARRELLI S.R.L",
		},
		{
			Name: "VALENTE MASSIMO",
		},
		{
			Name: "VALENTINA GIOVAGNINI ONLUS",
		},
		{
			Name: "VALENTINI PAOLO",
		},
		{
			Name: "VALERI ANGELO",
		},
		{
			Name: "VALERI GIAMPAOLO",
		},
		{
			Name: "VALLINI ILARIA",
		},
		{
			Name: "VANACORE GIULIO",
		},
		{
			Name: "VECCHIA CANTINA MONTEPULCIANO",
		},
		{
			Name: "VECOM CORNICI DI BELIGNI SANTI",
		},
		{
			Name: "VEIZAJ MANJOLA  ANTICO FORNO",
		},
		{
			Name: "VELIER S.P.A",
		},
		{
			Name: "VELVET GARDEN S.R.L.",
		},
		{
			Name: "VELVET UNDERGROUND DI",
		},
		{
			Name: "VENDITA AL MINUTO (CASSA)",
		},
		{
			Name: "VERSARI ANTONELLA",
		},
		{
			Name: "VETRERIA CASTIGLIONESE S.R.L.",
		},
		{
			Name: "VIA GUELFA 74 DI VALENTINA SUSSI",
		},
		{
			Name: "VIA GUIDO MONACO S.R.L.",
		},
		{
			Name: "VIA VITTORIO VENETO S.R.L.",
		},
		{
			Name: "VIBRO MANUFATTI DI GIANNINI",
		},
		{
			Name: "VICOLO DELL'AMORINO DI",
		},
		{
			Name: "VIGILANZA ANTINCENDIO BOSCHIVO",
		},
		{
			Name: "VILLA IL CIPRESSO SOCIETA",
		},
		{
			Name: "VILLA LOGGIO S.R.L.",
		},
		{
			Name: "VILLA MARSILI S.P.A",
		},
		{
			Name: "VILLAGGIO DEL GIOVANE",
		},
		{
			Name: "VINARIUM DI ROSSI LUCIANO",
		},
		{
			Name: "VINERIA DEL CORSO DI AREZZO",
		},
		{
			Name: "VINTAGE CLUB",
		},
		{
			Name: "VIRTUS BUONCONVENTO SOCIETA",
		},
		{
			Name: "VITI GIOVANNI BAR TABACCHI",
		},
		{
			Name: "VITTORIO ARGENTO DI",
		},
		{
			Name: "VIVAI FRAPPI MARCO",
		},
		{
			Name: "WAHEED ABDUL BD ABDUL",
		},
		{
			Name: "WAQAE ASHRAF JAMUNA ALIMENTARI",
		},
		{
			Name: "WHITE ROSE S.A.S.",
		},
		{
			Name: "WIMCO S.R.L.",
		},
		{
			Name: "WWW.LA PIZZA.DI GEPPONI",
		},
		{
			Name: "WWW.ORPC.IT S.R.L.",
		},
		{
			Name: "YO.YO SELF DI CROCCOLINO",
		},
		{
			Name: "ZADA MUHAMMAU",
		},
		{
			Name: "ZALL KEBAB 1 DI KORMAZ MEHMET",
		},
		{
			Name: "ZAM ZAM KEBAB DI REHMAN",
		},
		{
			Name: "ZETA DUE S.N.C.",
		},
		{
			Name: "ZINGALE EMANUELE",
		},
		{
			Name: "ZOPPAS FEDERICO",
		},
		{
			Name: "ZORZI 2.0 S.R.L.",
		},
		{
			Name: "ZUEGG S.P.A.",
		},
		{
			Name: "ZUFOLI GIAMPIERO TABACCHI N 10",
		},
		{
			Name: "CONTI MARZIA E PATRIZIA SNC",
		},
		{
			Name: "ROSALIA SNC BAR ROSALIA",
		},
		{
			Name: "GISTORE S.R.L.",
		},
		{
			Name: "BEN-BAR SAS DI RONDUCCI",
		},
		{
			Name: "CHICKEN HOUSE S.N.C.",
		},
		{
			Name: "BEGUM SHAHNAZ MALIK PIZZA",
		},
		{
			Name: "HASU SUSCHI",
		},
		{
			Name: "VANELLI S.R.L.",
		},
		{
			Name: "MARKET LA  BOTTEGA DI PACIFICO",
		},
		{
			Name: "ITALGIAD SRL",
		},
		{
			Name: "SER GIAN  DI CANDUSSI SERENA",
		},
		{
			Name: "SCAPECCHI FABIO GELATERIA GOLOGIO",
		},
		{
			Name: "OKONTA BUAKU ",
		},
		{
			Name: "MACRI COSMA & FIGLI S.R.L.",
		},
		{
			Name: "GIDUE  S.R.L.",
		},
		{
			Name: "RISTORANTE KILOMETRO ZERO DI GIRONI ROBERTO",
		},
		{
			Name: "IL CERVO DI PERUZZI ANDREA E PERUZZI VITTORIO",
		},
		{
			Name: "MORETTI DANIELA E MARCO",
		},
		{
			Name: "IL CIPRESSO DI MARCELLI ALBERTO",
		},
		{
			Name: "TREEMME S.P.A.",
		},
		{
			Name: "LA BOUTIQUE DEL PANE DI CHIARA  MERLUTTI",
		},
		{
			Name: "TASTY SRL",
		},
		{
			Name: "BARNESCHI MAURO ",
		},
		{
			Name: "UNICOM S.R.L.S.",
		},
		{
			Name: "CIRCOLO RICREATIVO ALBEROLO",
		},
		{
			Name: "ORIANA FRUTTA DI MENCARELLI SIMONE",
		},
		{
			Name: "PERE E MARGHERITE S.R.L.",
		},
		{
			Name: "VIA BOET S.N.C. DI BERENDSEN BEOT",
		},
		{
			Name: "BAR VITTORIA DI DURANTE",
		},
		{
			Name: "PUNTO GELATO GOVERNINI MARCO",
		},
		{
			Name: "BAR SPORT DI FRANCESCO CHEN JAN LONG",
		},
		{
			Name: "LA CASINA DEL PRATOS.A.S",
		},
		{
			Name: "SKY BAR LOUNGE S.N.C. DI GOMES",
		},
		{
			Name: "MASSIMO PIOVOSI MO PIZZA",
		},
		{
			Name: "LA MAGGIORE S.N.C. DI SILVI E. RINCHI G.",
		},
		{
			Name: "EUROPEA SERVIZI S.R.L.",
		},
		{
			Name: "MACELLERIA RA.RE. BONTA",
		},
		{
			Name: "LAHO HYSEN",
		},
		{
			Name: "CE.D.A.S. DI SALVINI EMANUELA",
		},
		{
			Name: "VESTRI CIOCCOLATO S.R.L.",
		},
		{
			Name: "VALDIVINO DI LAHO HYSEN",
		},
		{
			Name: "SOCIETA  AGRICOLA G.M.M. S.N.C.",
		},
		{
			Name: "COMUNE DI CIVITELLA IN VAL DI CHIANA",
		},
	}
	return data.CreateCustomerList(customerList)
}
