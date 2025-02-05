package activedirectory

import "github.com/lkarlslund/adalanche/modules/engine"

var (
	EdgeACLContainsDeny = engine.NewEdge("ACLContainsDeny").RegisterProbabilityCalculator(func(source, target *engine.Object) engine.Probability { return -1 })
	EdgeResetPassword   = engine.NewEdge("ResetPassword").RegisterProbabilityCalculator(func(source, target *engine.Object) engine.Probability {
		if uac, ok := target.AttrInt(UserAccountControl); ok && uac&engine.UAC_ACCOUNTDISABLE != 0 {
			return -1
		}
		return 100
	})
	EdgeReadPasswordId = engine.NewEdge("ReadPasswordId").SetDefault(false, false, false).RegisterProbabilityCalculator(func(source, target *engine.Object) engine.Probability {
		return 5
	})
	EdgeOwns             = engine.NewEdge("Owns")
	EdgeGenericAll       = engine.NewEdge("GenericAll")
	EdgeWriteAll         = engine.NewEdge("WriteAll")
	EdgeWritePropertyAll = engine.NewEdge("WritePropertyAll")
	EdgeWriteExtendedAll = engine.NewEdge("WriteExtendedAll")
	EdgeTakeOwnership    = engine.NewEdge("TakeOwnership")
	EdgeWriteDACL        = engine.NewEdge("WriteDACL")
	EdgeWriteSPN         = engine.NewEdge("WriteSPN").RegisterProbabilityCalculator(func(source, target *engine.Object) engine.Probability {
		if uac, ok := target.AttrInt(UserAccountControl); ok && uac&0x0002 /*UAC_ACCOUNTDISABLE*/ != 0 {
			// Account is disabled
			return 0
		}
		return 50
	})
	EdgeWriteValidatedSPN = engine.NewEdge("WriteValidatedSPN").RegisterProbabilityCalculator(func(source, target *engine.Object) engine.Probability {
		if uac, ok := target.AttrInt(UserAccountControl); ok && uac&0x0002 /*UAC_ACCOUNTDISABLE*/ != 0 {
			// Account is disabled
			return 0
		}
		return 50
	})
	EdgeWriteAllowedToAct       = engine.NewEdge("WriteAllowedToAct")
	EdgeAddMember               = engine.NewEdge("AddMember")
	EdgeAddMemberGroupAttr      = engine.NewEdge("AddMemberGroupAttr")
	EdgeAddSelfMember           = engine.NewEdge("AddSelfMember")
	EdgeReadMSAPassword         = engine.NewEdge("ReadMSAPassword")
	EdgeHasMSA                  = engine.NewEdge("HasMSA")
	EdgeWriteUserAccountControl = engine.NewEdge("WriteUserAccountControl").Describe("Allows attacker to set ENABLE and set DONT_REQ_PREAUTH and then to AS_REP Kerberoasting").RegisterProbabilityCalculator(func(source, target *engine.Object) engine.Probability {
		/*if uac, ok := target.AttrInt(activedirectory.UserAccountControl); ok && uac&0x0002 != 0 { //UAC_ACCOUNTDISABLE
			// Account is disabled
			return 0
		}*/
		return 50
	})

	EdgeWriteKeyCredentialLink = engine.NewEdge("WriteKeyCredentialLink").RegisterProbabilityCalculator(func(source, target *engine.Object) engine.Probability {
		if uac, ok := target.AttrInt(UserAccountControl); ok && uac&0x0002 /*UAC_ACCOUNTDISABLE*/ != 0 {
			// Account is disabled
			var canenable bool
			source.Edges(engine.Out).Range(func(key *engine.Object, value engine.EdgeBitmap) bool {
				if key == target {
					if value.IsSet(EdgeWriteUserAccountControl) {
						canenable = true
						return false
					}
				}
				return true
			})
			if !canenable {
				return 0
			}
		}
		return 100
	})
	EdgeWriteAttributeSecurityGUID           = engine.NewEdge("WriteAttrSecurityGUID").RegisterProbabilityCalculator(func(source, target *engine.Object) engine.Probability { return 5 }) // Only if you patch the DC, so this will actually never work
	EdgeSIDHistoryEquality                   = engine.NewEdge("SIDHistoryEquality")
	EdgeAllExtendedRights                    = engine.NewEdge("AllExtendedRights")
	EdgeDSReplicationSyncronize              = engine.NewEdge("DSReplSync")
	EdgeDSReplicationGetChanges              = engine.NewEdge("DSReplGetChngs").SetDefault(false, false, false)
	EdgeDSReplicationGetChangesAll           = engine.NewEdge("DSReplGetChngsAll").SetDefault(false, false, false)
	EdgeDSReplicationGetChangesInFilteredSet = engine.NewEdge("DSReplGetChngsInFiltSet").SetDefault(false, false, false)
	EdgeDCsync                               = engine.NewEdge("DCsync")
	EdgeReadLAPSPassword                     = engine.NewEdge("ReadLAPSPassword")
	EdgeMemberOfGroup                        = engine.NewEdge("MemberOfGroup")
	EdgeMemberOfGroupIndirect                = engine.NewEdge("MemberOfGroupIndirect").SetDefault(false, false, false)
	EdgeHasSPN                               = engine.NewEdge("HasSPN").Describe("Kerberoastable by requesting Kerberos service ticket against SPN and then bruteforcing the ticket").RegisterProbabilityCalculator(func(source, target *engine.Object) engine.Probability {
		if uac, ok := target.AttrInt(UserAccountControl); ok && uac&0x0002 /*UAC_ACCOUNTDISABLE*/ != 0 {
			// Account is disabled
			return 0
		}
		return 50
	})
	EdgeDontReqPreauth = engine.NewEdge("DontReqPreauth").Describe("Kerberoastable by AS-REP by requesting a TGT and then bruteforcing the ticket").RegisterProbabilityCalculator(func(source, target *engine.Object) engine.Probability {
		if uac, ok := target.AttrInt(UserAccountControl); ok && uac&0x0002 /*UAC_ACCOUNTDISABLE*/ != 0 {
			// Account is disabled
			return 0
		}
		return 50
	})
	EdgeOverwritesACL              = engine.NewEdge("OverwritesACL")
	EdgeAffectedByGPO              = engine.NewEdge("AffectedByGPO")
	PartOfGPO                      = engine.NewEdge("PartOfGPO")
	EdgeLocalAdminRights           = engine.NewEdge("AdminRights")
	EdgeLocalRDPRights             = engine.NewEdge("RDPRights").RegisterProbabilityCalculator(func(source, target *engine.Object) engine.Probability { return 30 })
	EdgeLocalDCOMRights            = engine.NewEdge("DCOMRights").RegisterProbabilityCalculator(func(source, target *engine.Object) engine.Probability { return 50 })
	EdgeScheduledTaskOnUNCPath     = engine.NewEdge("SchedTaskOnUNCPath")
	EdgeMachineScript              = engine.NewEdge("MachineScript")
	EdgeWriteAltSecurityIdentities = engine.NewEdge("WriteAltSecIdent")
	EdgeWriteProfilePath           = engine.NewEdge("WriteProfilePath")
	EdgeWriteScriptPath            = engine.NewEdge("WriteScriptPath")
	EdgeCertificateEnroll          = engine.NewEdge("CertificateEnroll")
	EdgeCertificateAutoEnroll      = engine.NewEdge("CertificateAutoEnroll")
	EdgeVoodooBit                  = engine.NewEdge("VoodooBit")
)
