/**
 * OPERA Cloud Accounts Receivables API
 * APIs to cater for Accounts Receivables functionality in OPERA Cloud. <br /><br The OPERA Cloud Accounts Receivable module enables you to manage debtors’ accounts, invoices, and remittance.<br /><br /> Compatible with OPERA Cloud release 22.3.0.0.<br /><br /><p> This document and all content within is available under the Universal Permissive License v 1.0 (https://oss.oracle.com/licenses/upl). Copyright (c) 2020, 2022 Oracle and/or its affiliates.</p>
 *
 * The version of the OpenAPI document: 22.3.0.0
 * Contact: hospitality_apis_ww_grp@oracle.com
 *
 * NOTE: This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
 * https://openapi-generator.tech
 * Do not edit the class manually.
 */
import type { ARInvoiceCategory } from './ARInvoiceCategory';
import type { CashierInfoType } from './CashierInfoType';
import type { CurrencyAmountType } from './CurrencyAmountType';
import type { FiscalResponseType } from './FiscalResponseType';
import type { FolioStatusType } from './FolioStatusType';
import type { FolioTextsTypeInner } from './FolioTextsTypeInner';
import type { MarketCodeInfoType } from './MarketCodeInfoType';
import type { ProfileInfoType } from './ProfileInfoType';
import type { ResPaymentCardType } from './ResPaymentCardType';
import type { ReservationId } from './ReservationId';
import type { ReservationInfoType } from './ReservationInfoType';
import type { RoomClassCodeInfoType } from './RoomClassCodeInfoType';
import type { SourceCodeInfoType } from './SourceCodeInfoType';
import type { TrxInfoType } from './TrxInfoType';
import type { UniqueIDType } from './UniqueIDType';
/**
 * Details of an AR Invoice.
 * @export
 * @interface ARPaymentInvoicesTypeInner
 */
export interface ARPaymentInvoicesTypeInner {
    /**
     * Property where the invoice exists.
     * @type {string}
     * @memberof ARPaymentInvoicesTypeInner
     */
    hotelId?: string;
    /**
     *
     * @type {ReservationId}
     * @memberof ARPaymentInvoicesTypeInner
     */
    reservationId?: ReservationId;
    /**
     *
     * @type {UniqueIDType}
     * @memberof ARPaymentInvoicesTypeInner
     */
    guestProfileId?: UniqueIDType;
    /**
     * The Aging age of the invoice
     * @type {number}
     * @memberof ARPaymentInvoicesTypeInner
     */
    age?: number;
    /**
     * Name of the Guest who consumed these transactions.
     * @type {string}
     * @memberof ARPaymentInvoicesTypeInner
     */
    guestName?: string;
    /**
     *
     * @type {CurrencyAmountType}
     * @memberof ARPaymentInvoicesTypeInner
     */
    originalAmount?: CurrencyAmountType;
    /**
     *
     * @type {CurrencyAmountType}
     * @memberof ARPaymentInvoicesTypeInner
     */
    amount?: CurrencyAmountType;
    /**
     *
     * @type {CurrencyAmountType}
     * @memberof ARPaymentInvoicesTypeInner
     */
    payments?: CurrencyAmountType;
    /**
     *
     * @type {CurrencyAmountType}
     * @memberof ARPaymentInvoicesTypeInner
     */
    balance?: CurrencyAmountType;
    /**
     * User-defined posting reference.
     * @type {string}
     * @memberof ARPaymentInvoicesTypeInner
     */
    reference?: string;
    /**
     * User-defined posting remark.
     * @type {string}
     * @memberof ARPaymentInvoicesTypeInner
     */
    remark?: string;
    /**
     *
     * @type {boolean}
     * @memberof ARPaymentInvoicesTypeInner
     */
    adjusted?: boolean;
    /**
     *
     * @type {number}
     * @memberof ARPaymentInvoicesTypeInner
     */
    parentInvoiceNo?: number;
    /**
     *
     * @type {boolean}
     * @memberof ARPaymentInvoicesTypeInner
     */
    compressed?: boolean;
    /**
     *
     * @type {Date}
     * @memberof ARPaymentInvoicesTypeInner
     */
    compressedDate?: Date;
    /**
     *
     * @type {boolean}
     * @memberof ARPaymentInvoicesTypeInner
     */
    transferredOut?: boolean;
    /**
     *
     * @type {boolean}
     * @memberof ARPaymentInvoicesTypeInner
     */
    transferredIn?: boolean;
    /**
     *
     * @type {MarketCodeInfoType}
     * @memberof ARPaymentInvoicesTypeInner
     */
    market?: MarketCodeInfoType;
    /**
     *
     * @type {RoomClassCodeInfoType}
     * @memberof ARPaymentInvoicesTypeInner
     */
    roomClass?: RoomClassCodeInfoType;
    /**
     *
     * @type {SourceCodeInfoType}
     * @memberof ARPaymentInvoicesTypeInner
     */
    source?: SourceCodeInfoType;
    /**
     *
     * @type {ReservationInfoType}
     * @memberof ARPaymentInvoicesTypeInner
     */
    reservationInfo?: ReservationInfoType;
    /**
     *
     * @type {CashierInfoType}
     * @memberof ARPaymentInvoicesTypeInner
     */
    cashierInfo?: CashierInfoType;
    /**
     *
     * @type {TrxInfoType}
     * @memberof ARPaymentInvoicesTypeInner
     */
    transactionInfo?: TrxInfoType;
    /**
     *
     * @type {ResPaymentCardType}
     * @memberof ARPaymentInvoicesTypeInner
     */
    paymentCard?: ResPaymentCardType;
    /**
     *
     * @type {ProfileInfoType}
     * @memberof ARPaymentInvoicesTypeInner
     */
    addresseeInfo?: ProfileInfoType;
    /**
     * Flag to check Partail Transfer Allowed.
     * @type {boolean}
     * @memberof ARPaymentInvoicesTypeInner
     */
    allowPartialTransferYn?: boolean;
    /**
     * Flag to check Invoice Statement is Printed.
     * @type {boolean}
     * @memberof ARPaymentInvoicesTypeInner
     */
    printed?: boolean;
    /**
     * Invoice Statement Printed Date.
     * @type {Date}
     * @memberof ARPaymentInvoicesTypeInner
     */
    printedDate?: Date;
    /**
     *
     * @type {UniqueIDType}
     * @memberof ARPaymentInvoicesTypeInner
     */
    storedFolioId?: UniqueIDType;
    /**
     * Name of the Stored Folio.
     * @type {string}
     * @memberof ARPaymentInvoicesTypeInner
     */
    storedFolioName?: string;
    /**
     *
     * @type {UniqueIDType}
     * @memberof ARPaymentInvoicesTypeInner
     */
    storedDebitFolioId?: UniqueIDType;
    /**
     * Name of the Stored Debit Folio.
     * @type {string}
     * @memberof ARPaymentInvoicesTypeInner
     */
    storedDebitFolioName?: string;
    /**
     * This stores the description for the type of tax calculation especially with tax exemption, etc.
     * @type {Array<FolioTextsTypeInner>}
     * @memberof ARPaymentInvoicesTypeInner
     */
    folioTexts?: Array<FolioTextsTypeInner>;
    /**
     * Unique Custom Numbers associated with this record.
     * @type {Array<string>}
     * @memberof ARPaymentInvoicesTypeInner
     */
    customNumbers?: Array<string>;
    /**
     *
     * @type {FiscalResponseType}
     * @memberof ARPaymentInvoicesTypeInner
     */
    fiscalResponseType?: FiscalResponseType;
    /**
     * Unique Transaction Identifier of the Invoice.
     * @type {number}
     * @memberof ARPaymentInvoicesTypeInner
     */
    transactionNo?: number;
    /**
     * Transaction Date of the invoice.
     * @type {Date}
     * @memberof ARPaymentInvoicesTypeInner
     */
    transactionDate?: Date;
    /**
     * Status of the invoice.
     * @type {string}
     * @memberof ARPaymentInvoicesTypeInner
     */
    status?: string;
    /**
     * Transaction code of the invoice.
     * @type {string}
     * @memberof ARPaymentInvoicesTypeInner
     */
    transactionCode?: string;
    /**
     * Folio Number.
     * @type {number}
     * @memberof ARPaymentInvoicesTypeInner
     */
    folioNo?: number;
    /**
     * Invoice No after the folio is generated. Same invoice number may be referred in multiple folios
     * @type {number}
     * @memberof ARPaymentInvoicesTypeInner
     */
    invoiceNo?: number;
    /**
     * The Fiscal Bill number of this posting
     * @type {string}
     * @memberof ARPaymentInvoicesTypeInner
     */
    fiscalBillNo?: string;
    /**
     * The name of the Folio Type used for the Folio Number sequence.
     * @type {string}
     * @memberof ARPaymentInvoicesTypeInner
     */
    folioTypeName?: string;
    /**
     * Internal window ID which is unique to the reservation. This ID can only be used for reference.
     * @type {string}
     * @memberof ARPaymentInvoicesTypeInner
     */
    internalFolioWindowID?: string;
    /**
     * Date of Folio Generation.
     * @type {Date}
     * @memberof ARPaymentInvoicesTypeInner
     */
    folioDate?: Date;
    /**
     *
     * @type {FolioStatusType}
     * @memberof ARPaymentInvoicesTypeInner
     */
    folioStatus?: FolioStatusType;
    /**
     * The folio number with prefix value.
     * @type {string}
     * @memberof ARPaymentInvoicesTypeInner
     */
    folioNoWithPrefix?: string;
    /**
     *
     * @type {ARInvoiceCategory}
     * @memberof ARPaymentInvoicesTypeInner
     */
    invoiceType?: ARInvoiceCategory;
    /**
     * Statementno of the invoice.
     * @type {number}
     * @memberof ARPaymentInvoicesTypeInner
     */
    statementNo?: number;
    /**
     * Revenue date ID of the invoice.
     * @type {Date}
     * @memberof ARPaymentInvoicesTypeInner
     */
    revenueDate?: Date;
    /**
     * Close date of the invoice.
     * @type {Date}
     * @memberof ARPaymentInvoicesTypeInner
     */
    closeDate?: Date;
    /**
     * Posting date of the invoice.
     * @type {Date}
     * @memberof ARPaymentInvoicesTypeInner
     */
    postingDate?: Date;
    /**
     * Transfer date of the invoice.
     * @type {Date}
     * @memberof ARPaymentInvoicesTypeInner
     */
    transferDate?: Date;
    /**
     * Flag to indicate if postings or adjustments can be made to the invoice.
     * @type {boolean}
     * @memberof ARPaymentInvoicesTypeInner
     */
    canBeModified?: boolean;
    /**
     *
     * @type {CurrencyAmountType}
     * @memberof ARPaymentInvoicesTypeInner
     */
    amountToApply?: CurrencyAmountType;
}
/**
 * Check if a given object implements the ARPaymentInvoicesTypeInner interface.
 */
export declare function instanceOfARPaymentInvoicesTypeInner(value: object): boolean;
export declare function ARPaymentInvoicesTypeInnerFromJSON(json: any): ARPaymentInvoicesTypeInner;
export declare function ARPaymentInvoicesTypeInnerFromJSONTyped(json: any, ignoreDiscriminator: boolean): ARPaymentInvoicesTypeInner;
export declare function ARPaymentInvoicesTypeInnerToJSON(value?: ARPaymentInvoicesTypeInner | null): any;
