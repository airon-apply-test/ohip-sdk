/**
 * OPERA Cloud Reservation API
 * APIs to cater for Reservation functionality in OPERA Cloud. <br /><br />OPERA Cloud Reservations provides a complete set of capabilities for creating and updating reservations. Reservations are a central feature of OPERA Cloud. As a key source of information, the reservation specifies a guest&apos;s arrival date, departure date, room type, rate, packages, and many other details. It is also a gateway to dozens of other functions that contribute to the guest&apos;s experience.  All reservations in OPERA Cloud require a guest profile.<br /><br /> You can create profiles while booking a reservation. If a profile already exists, you can look it up (using getProfiles in the Customer Relationship Management module) and attach it to the reservation during the reservation booking process using the profileId.<br /><br /> Compatible with OPERA Cloud release 21.5.0.0.<br /><br /><p> This document and all content within is available under the Universal Permissive License v 1.0 (https://oss.oracle.com/licenses/upl). Copyright (c) 2020, 2022 Oracle and/or its affiliates.</p>
 *
 * The version of the OpenAPI document: 21.5.0.0
 * Contact: hospitality_apis_ww_grp@oracle.com
 *
 * NOTE: This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
 * https://openapi-generator.tech
 * Do not edit the class manually.
 */
import type { AdvanceCheckInType } from './AdvanceCheckInType';
import type { AlertType } from './AlertType';
import type { AttachmentType } from './AttachmentType';
import type { AwardType } from './AwardType';
import type { CallType } from './CallType';
import type { CateringResInfoType } from './CateringResInfoType';
import type { CommentInfoType } from './CommentInfoType';
import type { CompAuthorizerInfoType } from './CompAuthorizerInfoType';
import type { ConfirmationType } from './ConfirmationType';
import type { CustomChargeExemptionType } from './CustomChargeExemptionType';
import type { ECouponType } from './ECouponType';
import type { ExternalReferenceType } from './ExternalReferenceType';
import type { FetchActivityBookingsType } from './FetchActivityBookingsType';
import type { FixedChargeType } from './FixedChargeType';
import type { GuestMessageType } from './GuestMessageType';
import type { HotelReservationTraceType } from './HotelReservationTraceType';
import type { HotelReservationTypeCancellation } from './HotelReservationTypeCancellation';
import type { HotelReservationTypeReservationProfiles } from './HotelReservationTypeReservationProfiles';
import type { HotelReservationTypeTransactionDiversions } from './HotelReservationTypeTransactionDiversions';
import type { HotelReservationsType } from './HotelReservationsType';
import type { IndicatorType } from './IndicatorType';
import type { LinkedReservationsInfoType } from './LinkedReservationsInfoType';
import type { MembershipType } from './MembershipType';
import type { NameValueDetailType } from './NameValueDetailType';
import type { OverrideInstructionType } from './OverrideInstructionType';
import type { PMSResStatusType } from './PMSResStatusType';
import type { PreferenceTypeType } from './PreferenceTypeType';
import type { PrepaidCardType } from './PrepaidCardType';
import type { ResAccessRestrictionType } from './ResAccessRestrictionType';
import type { ResCashieringType } from './ResCashieringType';
import type { ResCommunicationType } from './ResCommunicationType';
import type { ResGuestAdditionalInfoType } from './ResGuestAdditionalInfoType';
import type { ResGuestType } from './ResGuestType';
import type { ResHousekeepingType } from './ResHousekeepingType';
import type { ResInventoryItemsType } from './ResInventoryItemsType';
import type { ResSharedGuestInfoType } from './ResSharedGuestInfoType';
import type { ReservationAllowedActionType } from './ReservationAllowedActionType';
import type { ReservationECertificateType } from './ReservationECertificateType';
import type { ReservationLocatorType } from './ReservationLocatorType';
import type { ReservationMembershipType } from './ReservationMembershipType';
import type { ReservationPackageType } from './ReservationPackageType';
import type { ReservationPaymentMethodType } from './ReservationPaymentMethodType';
import type { ReservationPoliciesType } from './ReservationPoliciesType';
import type { ReservationQueueInformationType } from './ReservationQueueInformationType';
import type { RoomStayType } from './RoomStayType';
import type { RoutingInfoType } from './RoutingInfoType';
import type { ServiceRequest } from './ServiceRequest';
import type { SourceOfSaleType } from './SourceOfSaleType';
import type { TicketType } from './TicketType';
import type { TrackItItemType } from './TrackItItemType';
import type { UniqueIDType } from './UniqueIDType';
import type { UserDefinedFieldsType } from './UserDefinedFieldsType';
import type { WaitlistResType } from './WaitlistResType';
/**
 * The Reservation class contains the current reservation being created or altered.
 * @export
 * @interface HotelReservationType
 */
export interface HotelReservationType {
    /**
     * Unique Id that references an object uniquely in the system.
     * @type {Array<UniqueIDType>}
     * @memberof HotelReservationType
     */
    reservationIdList?: Array<UniqueIDType>;
    /**
     * This type contains unique information of external reference.
     * @type {Array<ExternalReferenceType>}
     * @memberof HotelReservationType
     */
    externalReferences?: Array<ExternalReferenceType>;
    /**
     *
     * @type {SourceOfSaleType}
     * @memberof HotelReservationType
     */
    sourceOfSale?: SourceOfSaleType;
    /**
     *
     * @type {RoomStayType}
     * @memberof HotelReservationType
     */
    roomStay?: RoomStayType;
    /**
     *
     * @type {CompAuthorizerInfoType}
     * @memberof HotelReservationType
     */
    compAuthorizer?: CompAuthorizerInfoType;
    /**
     * Collection of guests associated with the reservation.
     * @type {Array<ResGuestType>}
     * @memberof HotelReservationType
     */
    reservationGuests?: Array<ResGuestType>;
    /**
     * Collection of shared guest reservations.
     * @type {Array<ResSharedGuestInfoType>}
     * @memberof HotelReservationType
     */
    sharedGuests?: Array<ResSharedGuestInfoType>;
    /**
     *
     * @type {ResGuestAdditionalInfoType}
     * @memberof HotelReservationType
     */
    additionalGuestInfo?: ResGuestAdditionalInfoType;
    /**
     *
     * @type {HotelReservationTypeReservationProfiles}
     * @memberof HotelReservationType
     */
    reservationProfiles?: HotelReservationTypeReservationProfiles;
    /**
     *
     * @type {ResCommunicationType}
     * @memberof HotelReservationType
     */
    reservationCommunication?: ResCommunicationType;
    /**
     *
     * @type {Array<TrackItItemType>}
     * @memberof HotelReservationType
     */
    trackItItems?: Array<TrackItItemType>;
    /**
     *
     * @type {HotelReservationsType}
     * @memberof HotelReservationType
     */
    shares?: HotelReservationsType;
    /**
     * Holds the Locator Information.
     * @type {Array<ReservationLocatorType>}
     * @memberof HotelReservationType
     */
    locators?: Array<ReservationLocatorType>;
    /**
     * Attachment List.
     * @type {Array<AttachmentType>}
     * @memberof HotelReservationType
     */
    attachments?: Array<AttachmentType>;
    /**
     * Attachment List.
     * @type {Array<AttachmentType>}
     * @memberof HotelReservationType
     */
    webRegistrationCards?: Array<AttachmentType>;
    /**
     * Service request.
     * @type {Array<ServiceRequest>}
     * @memberof HotelReservationType
     */
    serviceRequests?: Array<ServiceRequest>;
    /**
     * Criteria for fetching Activity Bookings.
     * @type {Array<FetchActivityBookingsType>}
     * @memberof HotelReservationType
     */
    reservationActivities?: Array<FetchActivityBookingsType>;
    /**
     * Criteria for fetching Activity Bookings.
     * @type {Array<FetchActivityBookingsType>}
     * @memberof HotelReservationType
     */
    scheduledActivities?: Array<FetchActivityBookingsType>;
    /**
     * Holds fixed charge detail.
     * @type {Array<PrepaidCardType>}
     * @memberof HotelReservationType
     */
    prepaidCards?: Array<PrepaidCardType>;
    /**
     * List of awards.
     * @type {Array<AwardType>}
     * @memberof HotelReservationType
     */
    profileAwards?: Array<AwardType>;
    /**
     * A Service object represents a non-room product provided to guests. Service products may have associated inventory and charges.
     * @type {Array<ReservationPackageType>}
     * @memberof HotelReservationType
     */
    reservationPackages?: Array<ReservationPackageType>;
    /**
     *
     * @type {ResInventoryItemsType}
     * @memberof HotelReservationType
     */
    inventoryItems?: ResInventoryItemsType;
    /**
     * List of Notes for the Guest related to the reservation.
     * @type {Array<CommentInfoType>}
     * @memberof HotelReservationType
     */
    comments?: Array<CommentInfoType>;
    /**
     * List of Notes for the Guest related to the profile.
     * @type {Array<CommentInfoType>}
     * @memberof HotelReservationType
     */
    guestComments?: Array<CommentInfoType>;
    /**
     * List of Memberships for the Guest related to the reservation.
     * @type {Array<MembershipType>}
     * @memberof HotelReservationType
     */
    guestMemberships?: Array<MembershipType>;
    /**
     * Collection of reservation preferences.
     * @type {Array<PreferenceTypeType>}
     * @memberof HotelReservationType
     */
    preferenceCollection?: Array<PreferenceTypeType>;
    /**
     * List of reservation memberships.
     * @type {Array<ReservationMembershipType>}
     * @memberof HotelReservationType
     */
    reservationMemberships?: Array<ReservationMembershipType>;
    /**
     * Defines reservation payment methods.
     * @type {Array<ReservationPaymentMethodType>}
     * @memberof HotelReservationType
     */
    reservationPaymentMethods?: Array<ReservationPaymentMethodType>;
    /**
     * A routing info object can either be of type Folio OR of type Room with its corresponding instructions.
     * @type {Array<RoutingInfoType>}
     * @memberof HotelReservationType
     */
    routingInstructions?: Array<RoutingInfoType>;
    /**
     *
     * @type {ReservationPoliciesType}
     * @memberof HotelReservationType
     */
    reservationPolicies?: ReservationPoliciesType;
    /**
     *
     * @type {ResCashieringType}
     * @memberof HotelReservationType
     */
    cashiering?: ResCashieringType;
    /**
     *
     * @type {ResHousekeepingType}
     * @memberof HotelReservationType
     */
    housekeeping?: ResHousekeepingType;
    /**
     *
     * @type {LinkedReservationsInfoType}
     * @memberof HotelReservationType
     */
    linkedReservation?: LinkedReservationsInfoType;
    /**
     *
     * @type {boolean}
     * @memberof HotelReservationType
     */
    extSystemSync?: boolean;
    /**
     *
     * @type {UserDefinedFieldsType}
     * @memberof HotelReservationType
     */
    userDefinedFields?: UserDefinedFieldsType;
    /**
     * Collection of lamp indicators.
     * @type {Array<IndicatorType>}
     * @memberof HotelReservationType
     */
    reservationIndicators?: Array<IndicatorType>;
    /**
     *
     * @type {WaitlistResType}
     * @memberof HotelReservationType
     */
    waitlist?: WaitlistResType;
    /**
     *
     * @type {HotelReservationTypeCancellation}
     * @memberof HotelReservationType
     */
    cancellation?: HotelReservationTypeCancellation;
    /**
     *
     * @type {CateringResInfoType}
     * @memberof HotelReservationType
     */
    catering?: CateringResInfoType;
    /**
     * List of alerts.
     * @type {Array<AlertType>}
     * @memberof HotelReservationType
     */
    alerts?: Array<AlertType>;
    /**
     * List of Reservation Traces.
     * @type {Array<HotelReservationTraceType>}
     * @memberof HotelReservationType
     */
    traces?: Array<HotelReservationTraceType>;
    /**
     * List of confirmation letters.
     * @type {Array<ConfirmationType>}
     * @memberof HotelReservationType
     */
    confirmationLetters?: Array<ConfirmationType>;
    /**
     * Holds call information.
     * @type {Array<CallType>}
     * @memberof HotelReservationType
     */
    callHistory?: Array<CallType>;
    /**
     * Holds fixed charge detail.
     * @type {Array<FixedChargeType>}
     * @memberof HotelReservationType
     */
    fixedCharges?: Array<FixedChargeType>;
    /**
     * Holds the Message Information
     * @type {Array<GuestMessageType>}
     * @memberof HotelReservationType
     */
    guestMessages?: Array<GuestMessageType>;
    /**
     * An id representing the record lock for this reservation. When passed, update operation will verify that this lock is still valid.
     * @type {number}
     * @memberof HotelReservationType
     */
    lockHandle?: number;
    /**
     * List of override actions done.
     * @type {Array<OverrideInstructionType>}
     * @memberof HotelReservationType
     */
    overrideInstructions?: Array<OverrideInstructionType>;
    /**
     *
     * @type {ReservationQueueInformationType}
     * @memberof HotelReservationType
     */
    queue?: ReservationQueueInformationType;
    /**
     * Allowed action.
     * @type {Array<ReservationAllowedActionType>}
     * @memberof HotelReservationType
     */
    allowedActions?: Array<ReservationAllowedActionType>;
    /**
     * Collection of ECouponType object.
     * @type {Array<ECouponType>}
     * @memberof HotelReservationType
     */
    eCoupons?: Array<ECouponType>;
    /**
     *
     * @type {HotelReservationTypeTransactionDiversions}
     * @memberof HotelReservationType
     */
    transactionDiversions?: HotelReservationTypeTransactionDiversions;
    /**
     *
     * @type {AdvanceCheckInType}
     * @memberof HotelReservationType
     */
    advanceCheckIn?: AdvanceCheckInType;
    /**
     * Collection of TicketType objects.
     * @type {Array<TicketType>}
     * @memberof HotelReservationType
     */
    tickets?: Array<TicketType>;
    /**
     * Configurable Access Restriction Message for Excluded Block or Rate Code.
     * @type {string}
     * @memberof HotelReservationType
     */
    accessRestrictionMessage?: string;
    /**
     * Collection of ECertificateType object.
     * @type {Array<ReservationECertificateType>}
     * @memberof HotelReservationType
     */
    eCertificates?: Array<ReservationECertificateType>;
    /**
     *
     * @type {NameValueDetailType}
     * @memberof HotelReservationType
     */
    customNameValueDetail?: NameValueDetailType;
    /**
     * List of Custom Charge Exemptions.
     * @type {Array<CustomChargeExemptionType>}
     * @memberof HotelReservationType
     */
    customChargeExemptionDetails?: Array<CustomChargeExemptionType>;
    /**
     * This flag will force to borrow from house if borrow required for number of rooms requested, this will only applicable when reservation is attached to a Block.
     * @type {boolean}
     * @memberof HotelReservationType
     */
    autoBorrowFromHouse?: boolean;
    /**
     * Forcefully override the create/update action on the Reservation. For real-time CRS interfaces, this element is used as force selling indicator in case of Create or Update Reservation
     * @type {boolean}
     * @memberof HotelReservationType
     */
    overrideExternalChecks?: boolean;
    /**
     *
     * @type {string}
     * @memberof HotelReservationType
     */
    hotelId?: string;
    /**
     * Boolean True if this reservation is reserving rooms. False if it is only reserving services.
     * @type {boolean}
     * @memberof HotelReservationType
     */
    roomStayReservation?: boolean;
    /**
     *
     * @type {PMSResStatusType}
     * @memberof HotelReservationType
     */
    reservationStatus?: PMSResStatusType;
    /**
     *
     * @type {PMSResStatusType}
     * @memberof HotelReservationType
     */
    computedReservationStatus?: PMSResStatusType;
    /**
     * When true, indicates the reservation is for a guest that walks-in without a reservation. When false, the reservation is not a walk-in.
     * @type {boolean}
     * @memberof HotelReservationType
     */
    walkIn?: boolean;
    /**
     * Indicates whether to show the rate or not on a printed document which includes the rate.
     * @type {boolean}
     * @memberof HotelReservationType
     */
    printRate?: boolean;
    /**
     * Time stamp of the creation.
     * @type {Date}
     * @memberof HotelReservationType
     */
    createDateTime?: Date;
    /**
     * ID of creator. The creator could be a software system identifier or an identifier of an employee resposible for the creation.
     * @type {string}
     * @memberof HotelReservationType
     */
    creatorId?: string;
    /**
     * Time stamp of last modification.
     * @type {Date}
     * @memberof HotelReservationType
     */
    lastModifyDateTime?: Date;
    /**
     * Identifies the last software system or person to modify a record.
     * @type {string}
     * @memberof HotelReservationType
     */
    lastModifierId?: string;
    /**
     * Date an item will be purged from a database (e.g., from a live database to an archive).
     * @type {Date}
     * @memberof HotelReservationType
     */
    purgeDate?: Date;
    /**
     * Business Date when the reservation was created.
     * @type {Date}
     * @memberof HotelReservationType
     */
    createBusinessDate?: Date;
    /**
     * Business Date when the reservation was last reinstated.
     * @type {Date}
     * @memberof HotelReservationType
     */
    reinstateDate?: Date;
    /**
     * Party/PartyCode when reservation is part of a party(a small group).
     * @type {string}
     * @memberof HotelReservationType
     */
    party?: string;
    /**
     * Name Of the Resort/Hotel for the primary enrollment
     * @type {string}
     * @memberof HotelReservationType
     */
    primaryEnrollmentResort?: string;
    /**
     * Name of the Chain for the primary enrollment
     * @type {string}
     * @memberof HotelReservationType
     */
    primaryEnrollmentChain?: string;
    /**
     * Custom Reference to identify reservation.
     * @type {string}
     * @memberof HotelReservationType
     */
    customReference?: string;
    /**
     * The color to display this reservation on room grid displays.
     * @type {string}
     * @memberof HotelReservationType
     */
    displayColor?: string;
    /**
     * Mark this reservation as recently accessed.
     * @type {boolean}
     * @memberof HotelReservationType
     */
    markAsRecentlyAccessed?: boolean;
    /**
     * Indicates to Override the Room Inventory Check if applicable.
     * @type {boolean}
     * @memberof HotelReservationType
     */
    overrideInventoryCheck?: boolean;
    /**
     *
     * @type {ResAccessRestrictionType}
     * @memberof HotelReservationType
     */
    accessRestriction?: ResAccessRestrictionType;
    /**
     * Is the reservation pre-registered.
     * @type {boolean}
     * @memberof HotelReservationType
     */
    preRegistered?: boolean;
    /**
     * Indicates whether the reservation is eligible for upgrade.
     * @type {boolean}
     * @memberof HotelReservationType
     */
    upgradeEligible?: boolean;
    /**
     *
     * @type {boolean}
     * @memberof HotelReservationType
     */
    overrideBlockRestriction?: boolean;
    /**
     * Flag to allow auto checkin
     * @type {boolean}
     * @memberof HotelReservationType
     */
    allowAutoCheckin?: boolean;
    /**
     *
     * @type {boolean}
     * @memberof HotelReservationType
     */
    hasOpenFolio?: boolean;
    /**
     * Flag containing true or false value for reservation to be eligible for self-checkout by guest using mobile device . Pass the 'true' or 'false' values when creating / modifying reservation to indicate whether a reservation is eligible for mobile checkout yes / no. Upon fetch, the current state of the flag will show true or false.
     * @type {boolean}
     * @memberof HotelReservationType
     */
    allowMobileCheckout?: boolean;
    /**
     * Attribute AllowMobileViewFolio is set to true when the reservation is eligible for viewing folio using mobile device.
     * @type {boolean}
     * @memberof HotelReservationType
     */
    allowMobileViewFolio?: boolean;
    /**
     * Flag containing true or false value for reservation to be eligible for prior check in by guest using mobile device . Pass the 'true' or 'false' values when creating / modifying reservation to indicate whether a reservation is eligible for mobile check in yes / no. Upon fetch, the current state of the flag will show true or false.
     * @type {boolean}
     * @memberof HotelReservationType
     */
    allowPreRegistration?: boolean;
    /**
     * Flag containing true or false value to determine if the guest has opted for communications pertaining to this reservation.Pass the 'true' or 'false' values when creating / modifying reservation to indicate whether a reservation is marked for communication. Upon fetch, the current state of the flag will show true or false.
     * @type {boolean}
     * @memberof HotelReservationType
     */
    optedForCommunication?: boolean;
    /**
     * This attribute represents Hotel charge card number. This number is generated automatically in OPERA when a reservation is created in OPERA. It is part of a functionality that provide the guest with the convenience of placing all their charges directly on their reservation folio.
     * @type {string}
     * @memberof HotelReservationType
     */
    chargeCardNumber?: string;
}
/**
 * Check if a given object implements the HotelReservationType interface.
 */
export declare function instanceOfHotelReservationType(value: object): boolean;
export declare function HotelReservationTypeFromJSON(json: any): HotelReservationType;
export declare function HotelReservationTypeFromJSONTyped(json: any, ignoreDiscriminator: boolean): HotelReservationType;
export declare function HotelReservationTypeToJSON(value?: HotelReservationType | null): any;
