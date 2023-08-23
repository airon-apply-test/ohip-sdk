/**
 * OPERA Cloud Block API
 * APIs to cater for Business Block functionality in OPERA Cloud. <br /><br /> A block is a group of rooms held for guests who are attending an event, meeting, or function. You can create blocks for family reunions, business conferences, weddings, and so on. You can also set aside rooms for the event (block).<br /><br /> Compatible with OPERA Cloud release 22.3.0.0.<br /><br /><p> This document and all content within is available under the Universal Permissive License v 1.0 (https://oss.oracle.com/licenses/upl). Copyright (c) 2020, 2022 Oracle and/or its affiliates.</p>
 *
 * The version of the OpenAPI document: 22.3.0.0
 * Contact: hospitality_apis_ww_grp@oracle.com
 *
 * NOTE: This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
 * https://openapi-generator.tech
 * Do not edit the class manually.
 */
/**
 * Supported housekeeping statistical codes.
 * @export
 */
export declare const HSKStatCodeType: {
    readonly ApprovedCompPostings: "ApprovedCompPostings";
    readonly StagedCompPostings: "StagedCompPostings";
    readonly DeclinedCompPostings: "DeclinedCompPostings";
    readonly CompRoutingInstructionsRequests: "CompRoutingInstructionsRequests";
    readonly TotalPhysicalRooms: "TotalPhysicalRooms";
    readonly TotalRoomsToSell: "TotalRoomsToSell";
    readonly TotalOutOfOrder: "TotalOutOfOrder";
    readonly TotalOutOfService: "TotalOutOfService";
    readonly TotalRevenue: "TotalRevenue";
    readonly PercentRoomsOccupied: "PercentRoomsOccupied";
    readonly StayoverRooms: "StayoverRooms";
    readonly StayoverPersons: "StayoverPersons";
    readonly StayoverVip: "StayoverVIP";
    readonly DeparturesExpectedRooms: "DeparturesExpectedRooms";
    readonly DeparturesExpectedPersons: "DeparturesExpectedPersons";
    readonly DeparturesExpectedVip: "DeparturesExpectedVIP";
    readonly DeparturesActualRooms: "DeparturesActualRooms";
    readonly DeparturesActualPersons: "DeparturesActualPersons";
    readonly DeparturesActualVip: "DeparturesActualVIP";
    readonly DeparturesInLastHour: "DeparturesInLastHour";
    readonly ArrivalsInLastHour: "ArrivalsInLastHour";
    readonly ArrivalsTotal: "ArrivalsTotal";
    readonly ArrivalsExpectedRooms: "ArrivalsExpectedRooms";
    readonly ArrivalsExpectedPersons: "ArrivalsExpectedPersons";
    readonly ArrivalsExpectedVip: "ArrivalsExpectedVIP";
    readonly ArrivalsExpectedRoomsMadeToday: "ArrivalsExpectedRoomsMadeToday";
    readonly ArrivalsExpectedPersonsMadeToday: "ArrivalsExpectedPersonsMadeToday";
    readonly ArrivalsExpectedVipMadeToday: "ArrivalsExpectedVIPMadeToday";
    readonly ArrivalsActualRooms: "ArrivalsActualRooms";
    readonly ArrivalsActualPersons: "ArrivalsActualPersons";
    readonly ArrivalsActualVip: "ArrivalsActualVIP";
    readonly ExtendedStaysRooms: "ExtendedStaysRooms";
    readonly ExtendedStaysPersons: "ExtendedStaysPersons";
    readonly ExtendedStaysVip: "ExtendedStaysVIP";
    readonly DeparturesTotal: "DeparturesTotal";
    readonly EarlyDeparturesRooms: "EarlyDeparturesRooms";
    readonly EarlyDeparturesPersons: "EarlyDeparturesPersons";
    readonly EarlyDeparturesVip: "EarlyDeparturesVIP";
    readonly DayUseRooms: "DayUseRooms";
    readonly DayUsePersons: "DayUsePersons";
    readonly DayUseVip: "DayUseVIP";
    readonly WalkInRooms: "WalkInRooms";
    readonly WalkInPersons: "WalkInPersons";
    readonly WalkInVip: "WalkInVIP";
    readonly CanceledOnArrivalRooms: "CanceledOnArrivalRooms";
    readonly CanceledOnArrivalPersons: "CanceledOnArrivalPersons";
    readonly CanceledOnArrivalVip: "CanceledOnArrivalVIP";
    readonly ComplimentaryArrivalRooms: "ComplimentaryArrivalRooms";
    readonly ComplimentaryArrivalPersons: "ComplimentaryArrivalPersons";
    readonly ComplimentaryArrivalVip: "ComplimentaryArrivalVIP";
    readonly ComplimentaryStayoverRooms: "ComplimentaryStayoverRooms";
    readonly ComplimentaryStayoverPersons: "ComplimentaryStayoverPersons";
    readonly ComplimentaryStayoverVip: "ComplimentaryStayoverVIP";
    readonly ComplimentaryDepartureRooms: "ComplimentaryDepartureRooms";
    readonly ComplimentaryDeparturePersons: "ComplimentaryDeparturePersons";
    readonly ComplimentaryDepartureVip: "ComplimentaryDepartureVIP";
    readonly HouseUseArrivalVip: "HouseUseArrivalVIP";
    readonly HouseUseStayoverRooms: "HouseUseStayoverRooms";
    readonly HouseUseStayoverPersons: "HouseUseStayoverPersons";
    readonly HouseUseStayoverVip: "HouseUseStayoverVIP";
    readonly HouseUseDepartureRooms: "HouseUseDepartureRooms";
    readonly HouseUseDeparturePersons: "HouseUseDeparturePersons";
    readonly HouseUseDepartureVip: "HouseUseDepartureVIP";
    readonly MinAvailableTonightRooms: "MinAvailableTonightRooms";
    readonly MaxOccupiedTonightRooms: "MaxOccupiedTonightRooms";
    readonly MaxOccupiedTonightPersons: "MaxOccupiedTonightPersons";
    readonly MaxOccupiedTonightVip: "MaxOccupiedTonightVIP";
    readonly MaxPercentageOccupiedTonightRooms: "MaxPercentageOccupiedTonightRooms";
    readonly BlocksNotPickedUp: "BlocksNotPickedUp";
    readonly IndividualRooms: "IndividualRooms";
    readonly IndividualPersons: "IndividualPersons";
    readonly IndividualVip: "IndividualVIP";
    readonly GroupAndBlockRooms: "GroupAndBlockRooms";
    readonly GroupAndBlockPersons: "GroupAndBlockPersons";
    readonly GroupAndBlockVip: "GroupAndBlockVIP";
    readonly RoomRevenue: "RoomRevenue";
    readonly AverageRoomRevenue: "AverageRoomRevenue";
    readonly InspectedRooms: "InspectedRooms";
    readonly InspectedVacant: "InspectedVacant";
    readonly InspectedAssigned: "InspectedAssigned";
    readonly InspectedOccupied: "InspectedOccupied";
    readonly CleanedRooms: "CleanedRooms";
    readonly CleanVacant: "CleanVacant";
    readonly CleanAssigned: "CleanAssigned";
    readonly CleanOccupied: "CleanOccupied";
    readonly DirtyVacant: "DirtyVacant";
    readonly DirtyAssigned: "DirtyAssigned";
    readonly DirtyOccupied: "DirtyOccupied";
    readonly PickupVacant: "PickupVacant";
    readonly PickupAssigned: "PickupAssigned";
    readonly PickupOccupied: "PickupOccupied";
    readonly OutOfOrderVacant: "OutOfOrderVacant";
    readonly OutOfOrderAssigned: "OutOfOrderAssigned";
    readonly OutOfOrderOccupied: "OutOfOrderOccupied";
    readonly OutOfServiceVacant: "OutOfServiceVacant";
    readonly OutOfServiceAssigned: "OutOfServiceAssigned";
    readonly OutOfServiceOccupied: "OutOfServiceOccupied";
    readonly QueueRooms: "QueueRooms";
    readonly TurndownRequired: "TurndownRequired";
    readonly TurndownNotRequired: "TurndownNotRequired";
    readonly TurndownCompletedRequired: "TurndownCompletedRequired";
    readonly AdultsInHouse: "AdultsInHouse";
    readonly ChildrenInHouse: "ChildrenInHouse";
    readonly AdultsExpectedCheckedOut: "AdultsExpectedCheckedOut";
    readonly ChildrenExpectedCheckedOut: "ChildrenExpectedCheckedOut";
    readonly AdultsDeparted: "AdultsDeparted";
    readonly ChildrenDeparted: "ChildrenDeparted";
    readonly InHouseRooms: "InHouseRooms";
    readonly InHouse: "InHouse";
    readonly MaxOccupancyPercentage: "MaxOccupancyPercentage";
    readonly Stayover: "Stayover";
    readonly TotalRoomsReserved: "TotalRoomsReserved";
    readonly RevPar: "RevPar";
    readonly SkipRooms: "SkipRooms";
    readonly SleepRooms: "SleepRooms";
    readonly HouseUseArrivalRooms: "HouseUseArrivalRooms";
    readonly HouseUseArrivalPersons: "HouseUseArrivalPersons";
    readonly AverageCheckInTime: "AverageCheckInTime";
    readonly CurrentAveWaitTime: "CurrentAveWaitTime";
    readonly CheckedInsTotal: "CheckedInsTotal";
    readonly ExpectedCheckInsTotal: "ExpectedCheckInsTotal";
    readonly CheckedOutsTotal: "CheckedOutsTotal";
    readonly ExpectedCheckOutsTotal: "ExpectedCheckOutsTotal";
    readonly ScheduledCheckOutsTotal: "ScheduledCheckOutsTotal";
    readonly RoomMaintenanceResolvedTotal: "RoomMaintenanceResolvedTotal";
    readonly RoomMaintenanceUnResolvedTotal: "RoomMaintenanceUnResolvedTotal";
    readonly PreRegisteredTotal: "PreRegisteredTotal";
    readonly VipPreRegisteredTotal: "VIPPreRegisteredTotal";
    readonly OpenFolioTotal: "OpenFolioTotal";
    readonly TurndownTotal: "TurndownTotal";
    readonly VipTurndownTotal: "VIPTurndownTotal";
    readonly VipGuestsArriving: "VIPGuestsArriving";
    readonly VipGuestsDeparting: "VIPGuestsDeparting";
    readonly VipGuestsTotal: "VIPGuestsTotal";
    readonly IndividualAdvanceCheckedInCurrent: "IndividualAdvanceCheckedInCurrent";
    readonly BlockAdvanceCheckedInCurrent: "BlockAdvanceCheckedInCurrent";
    readonly IndividualAdvanceCheckedInInhouse: "IndividualAdvanceCheckedInInhouse";
    readonly BlockAdvanceCheckedInInhouse: "BlockAdvanceCheckedInInhouse";
    readonly IndividualAdvanceCheckedInTotal: "IndividualAdvanceCheckedInTotal";
    readonly BlockAdvanceCheckedInTotal: "BlockAdvanceCheckedInTotal";
    readonly InHouseBlocksTotal: "InHouseBlocksTotal";
    readonly CancellationsTotal: "CancellationsTotal";
    readonly NewReservationsTotal: "NewReservationsTotal";
    readonly ArrivalResvs: "ArrivalResvs";
    readonly ArrivalVipResvs: "ArrivalVIPResvs";
    readonly ArrivalMemberResvs: "ArrivalMemberResvs";
    readonly ArrivalUnallocResvs: "ArrivalUnallocResvs";
    readonly ArrivalManualAssgnResvs: "ArrivalManualAssgnResvs";
    readonly ArrivalAiAssgnResvs: "ArrivalAIAssgnResvs";
    readonly ArrivalAiUpgResvs: "ArrivalAIUpgResvs";
    readonly ArrivalAiAssgnVipResvs: "ArrivalAIAssgnVIPResvs";
    readonly ArrivalAiAssgnMemberResvs: "ArrivalAIAssgnMemberResvs";
    readonly ArrivalAiAssgnOverridden: "ArrivalAIAssgnOverridden";
};
export type HSKStatCodeType = typeof HSKStatCodeType[keyof typeof HSKStatCodeType];
export declare function HSKStatCodeTypeFromJSON(json: any): HSKStatCodeType;
export declare function HSKStatCodeTypeFromJSONTyped(json: any, ignoreDiscriminator: boolean): HSKStatCodeType;
export declare function HSKStatCodeTypeToJSON(value?: HSKStatCodeType | null): any;
